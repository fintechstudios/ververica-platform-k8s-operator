package appmanager

import (
	"context"
	"fmt"
	appmanagerapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"
	vvperrors "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/errors"
	"os"
	"strings"
)

// TokenData wraps an API token
type TokenData struct {
	Name       string
	value      string
	wasCreated bool
}

// TokenManager handles the API token lifecycle.
type TokenManager interface {
	// TokenExists checks if a token exists in a namespace by name
	TokenExists(ctx context.Context, namespaceName, name string) (bool, error)
	// CreateToken creates a token under a namespace with a given role and returns the secret
	CreateToken(ctx context.Context, namespaceName, name, role string) (string, error)
	// RemoveToken deletes a token from a namespace and returns whether it existed
	RemoveToken(ctx context.Context, namespaceName, name string) (bool, error)
}

// AuthStore manages authorized contexts for AppManager API calls on a per-vpnamespace basis.
type AuthStore interface {
	ContextForNamespace(baseCtx context.Context, namespace string) (context.Context, error)
	RemoveAllCreatedTokens(ctx context.Context) ([]string, error)
}


const defaultTokenEnvVar = "APPMANAGER_API_TOKEN" // nolint:gosec

// One token per namespace
const tokenName = "vp-k8s-operator-admin-token" // nolint:gosec

type authStore struct {
	namespaceTokenCache map[string]*TokenData
	tokenManager        TokenManager
}

func newAuthStore(manager TokenManager) *authStore {
	return &authStore{
		namespaceTokenCache: make(map[string]*TokenData),
		tokenManager:        manager,
	}
}

func NewAuthStore(tokenManager TokenManager) AuthStore {
	return newAuthStore(tokenManager)
}

func (s *authStore) findTokenForNamespaceInEnv(namespaceName string) *string {
	namespaceTokenEnvVar := fmt.Sprintf("%s_%s", defaultTokenEnvVar, strings.ToUpper(namespaceName))

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
func (s *authStore) getTokenForNamespace(ctx context.Context, namespaceName string) (string, error) {
	if s.namespaceTokenCache[namespaceName] != nil {
		return s.namespaceTokenCache[namespaceName].value, nil
	}

	if token := s.findTokenForNamespaceInEnv(namespaceName); token != nil {
		s.namespaceTokenCache[namespaceName] = &TokenData{
			value:      *token,
			Name:       "Env Provided",
			wasCreated: false,
		}
		return *token, nil
	}

	var err error
	if tokenData, err := s.getOrCreateTokenForNamespace(ctx, namespaceName); err == nil {
		s.namespaceTokenCache[namespaceName] = tokenData
		return tokenData.value, nil
	}

	return "", err
}

// getOrCreateTokenForNamespace gets a token for a namespace from the token manager or creates one if none are found
func (s *authStore) getOrCreateTokenForNamespace(ctx context.Context, namespaceName string) (*TokenData, error) {
	exists, err := s.tokenManager.TokenExists(ctx, namespaceName, tokenName)
	if err != nil {
		return nil, err
	}

	if exists {
		// a token already exists -- must delete it as cannot re-fetch it after creation
		// TODO: perhaps we should implement a SecretStore backed by K8s secrets for created tokens
		if _, err = s.tokenManager.RemoveToken(ctx, namespaceName, tokenName); err != nil {
			return nil, err
		}
	}

	// we always want full-access tokens
	token, err := s.tokenManager.CreateToken(ctx, namespaceName, tokenName, "owner")
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
func (s *authStore) ContextForNamespace(baseCtx context.Context, namespaceName string) (context.Context, error) {
	var token string
	var err error
	if token, err = s.getTokenForNamespace(baseCtx, namespaceName); err != nil {
		return nil, vvperrors.WrapAuthContextError(namespaceName, err)
	}

	return context.WithValue(baseCtx, appmanagerapi.ContextAccessToken, token), nil
}

// RemoveAllCreatedTokens deletes all tokens that have been created by the store
func (s *authStore) RemoveAllCreatedTokens(ctx context.Context) ([]string, error) {
	var deletedTokens []string
	for namespace, tokenData := range s.namespaceTokenCache {
		existed, err := s.tokenManager.RemoveToken(ctx, namespace, tokenData.Name)
		if err != nil {
			return deletedTokens, err
		}
		if existed {
			deletedTokens = append(deletedTokens, namespace+"/"+tokenData.Name)
		}
	}
	return deletedTokens, nil
}

