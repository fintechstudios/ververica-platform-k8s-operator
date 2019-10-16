package appManager

import (
	"context"
	"fmt"
	"os"
	"strings"

	appManagerApi "github.com/fintechstudios/ververica-platform-k8s-controller/appmanager-api-client"
)

const DefaultTokenEnvVar = "VP_API_TOKEN"

type AuthStore struct {
	namespaceTokenCache map[string]*string
}

// AuthNotFoundError represents when no auth token can be found for a namespace
type AuthNotFoundError struct {
	Namespace string
}

func (err AuthNotFoundError) Error() string {
	return fmt.Sprintf("no VP API token found for namespace %s", err.Namespace)
}

func NewAuthStore() *AuthStore {
	return &AuthStore{
		namespaceTokenCache: make(map[string]*string),
	}
}

func (s *AuthStore) findTokenForNamespaceInEnv(namespace string) *string {
	namespaceTokenEnvVar := fmt.Sprintf("VP_API_TOKEN_%s", namespace)

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] == namespaceTokenEnvVar {
			return &pair[1]
		}
	}

	if token, exists := os.LookupEnv(DefaultTokenEnvVar); exists {
		return &token
	}

	return nil
}

func (s *AuthStore) getTokenForNamespace(namespace string) (string, error) {
	// search chain:
	// - memory cache
	// - environment in form VP_API_TOKEN_{NAMESPACE}
	// - environment in form VP_API_TOKEN
	capNamespace := strings.ToUpper(namespace)
	if s.namespaceTokenCache[capNamespace] != nil {
		return *s.namespaceTokenCache[capNamespace], nil
	}

	if token := s.findTokenForNamespaceInEnv(capNamespace); token != nil {
		s.namespaceTokenCache[capNamespace] = token
		return *token, nil
	}

	return "", AuthNotFoundError{Namespace: namespace}
}

func (s *AuthStore) ContextForNamespace(namespace string) (context.Context, error) {
	var apiToken string
	var err error
	if apiToken, err = s.getTokenForNamespace(namespace); err != nil {
		return nil, err
	}

	return context.WithValue(context.Background(), appManagerApi.ContextAccessToken, apiToken), nil
}
