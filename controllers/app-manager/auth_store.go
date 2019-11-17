package appManager

import (
	"context"
	"fmt"
	"os"
	"strings"

	appManagerApi "github.com/fintechstudios/ververica-platform-k8s-controller/appmanager-api-client"
)

const defaultTokenEnvVar = "APPMANAGER_API_TOKEN"

// One token per namespace
const tokenName = "vp-k8s-controller-admin-token"

type TokenData struct {
	Name       string
	value      string
	wasCreated bool
}

// TokenNotFoundError represents when no auth token can be found for a namespace
type TokenNotFoundError struct {
	Namespace string
	Name      string
}

func (err TokenNotFoundError) Error() string {
	return fmt.Sprintf("no API token by name %s found in namespace %s", err.Name, err.Namespace)
}

type TokenManager interface {
	// TokenExists checks if a token exists in a namespace by name
	TokenExists(ctx context.Context, name, namespace string) (bool, error)
	// CreateToken creates a token under a namespace with a given role and returns the secret
	CreateToken(ctx context.Context, name, role, namespace string) (string, error)
	// RemoveToken deletes a token from a namespace and returns whether it existed
	RemoveToken(ctx context.Context, name, namespace string) (bool, error)
}

type AuthStore struct {
	namespaceTokenCache map[string]*TokenData
	tokenManager        TokenManager
}

func NewAuthStore(tokenManager TokenManager) *AuthStore {
	return &AuthStore{
		namespaceTokenCache: make(map[string]*TokenData),
		tokenManager:        tokenManager,
	}
}

func (s *AuthStore) findTokenForNamespaceInEnv(namespace string) *string {
	namespaceTokenEnvVar := fmt.Sprintf("%s_%s", defaultTokenEnvVar, strings.ToUpper(namespace))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] == namespaceTokenEnvVar {
			return &pair[1]
		}
	}

	if token, exists := os.LookupEnv(defaultTokenEnvVar); exists {
		return &token
	}

	return nil
}

// getTokenForNamespace attempts to get an API token for authenticating with the AppManager API under a namespace
// search chain:
// - memory cache
// - environment in form VP_API_TOKEN_{NAMESPACE}
// - environment in form VP_API_TOKEN
// if none are found, will attempt to create a token
func (s *AuthStore) getTokenForNamespace(ctx context.Context, namespace string) (string, error) {
	if s.namespaceTokenCache[namespace] != nil {
		return s.namespaceTokenCache[namespace].value, nil
	}

	if token := s.findTokenForNamespaceInEnv(namespace); token != nil {
		s.namespaceTokenCache[namespace] = &TokenData{
			value:      *token,
			Name:       "Env Provided",
			wasCreated: false,
		}
		return *token, nil
	}

	var err error
	if tokenData, err := s.getOrCreateTokenForNamespace(ctx, namespace); err == nil {
		s.namespaceTokenCache[namespace] = tokenData
		return tokenData.value, nil
	}

	return "", err
}

// getOrCreateTokenForNamespace gets a token for a namespace from the token manager or creates one if none are found
func (s *AuthStore) getOrCreateTokenForNamespace(ctx context.Context, namespace string) (*TokenData, error) {
	exists, err := s.tokenManager.TokenExists(ctx, tokenName, namespace)
	if err != nil {
		return nil, err
	}

	if exists {
		// a token already exists -- must delete it as cannot re-fetch it after creation
		// TODO: perhaps we should implement a SecretStore backed by K8s secrets for created tokens
		if _, err = s.tokenManager.RemoveToken(ctx, tokenName, namespace); err != nil {
			return nil, err
		}
	}

	// we always want full-access tokens
	token, err := s.tokenManager.CreateToken(ctx, tokenName, "owner", namespace)
	if err != nil {
		return nil, err
	}

	return &TokenData{
		Name:       tokenName,
		value:      token,
		wasCreated: true,
	}, nil
}

// ContextForNamespace gets a context with an authorization token for a namespace
func (s *AuthStore) ContextForNamespace(baseCtx context.Context, namespace string) (context.Context, error) {
	var token string
	var err error
	if token, err = s.getTokenForNamespace(baseCtx, namespace); err != nil {
		return nil, err
	}

	return context.WithValue(baseCtx, appManagerApi.ContextAccessToken, token), nil
}

// RemoveAllCreatedTokens deletes all tokens that have been created by the store
func (s *AuthStore) RemoveAllCreatedTokens(ctx context.Context) ([]string, error) {
	var deletedTokens []string
	for namespace, tokenData := range s.namespaceTokenCache {
		existed, err := s.tokenManager.RemoveToken(ctx, tokenData.Name, namespace)
		if err != nil {
			return deletedTokens, err
		}
		if existed {
			deletedTokens = append(deletedTokens, namespace+"/"+tokenData.Name)
		}
	}
	return deletedTokens, nil
}
