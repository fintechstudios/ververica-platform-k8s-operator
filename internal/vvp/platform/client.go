package platform

import (
	"context"
	"fmt"
	vvperrors "github.com/fintechstudios/ververica-platform-k8s-operator/internal/vvp/errors"
	platformapi "github.com/fintechstudios/ververica-platform-k8s-operator/internal/vvp/platform-api-client"
	"regexp"
	"strings"
)

type Client interface {
	Namespaces() NamespacesService
	ApiTokens() ApiTokensService
}

type client struct {
	apiClient         *platformapi.APIClient
	namespacesService NamespacesService
	apiTokensService  ApiTokensService
}

func NewClient(apiClient *platformapi.APIClient) Client {
	return &client{
		apiClient: apiClient,
	}
}

func (c *client) Namespaces() NamespacesService {
	if c.namespacesService == nil {
		c.namespacesService = &namespacesService{client: c}
	}
	return c.namespacesService
}

func (c *client) ApiTokens() ApiTokensService {
	if c.apiTokensService == nil {
		c.apiTokensService = &apiTokensService{client: c}
	}
	return c.apiTokensService
}

// Namespaces

type NamespacesService interface {
	GetNamespace(ctx context.Context, namespaceName string) (*platformapi.Namespace, error)
	CreateNamespace(ctx context.Context, namespace platformapi.Namespace) (*platformapi.Namespace, error)
	UpdateNamespace(ctx context.Context, namespaceName string, namespace platformapi.Namespace) (*platformapi.Namespace, error)
	DeleteNamespace(ctx context.Context, namespaceName string) (*platformapi.Namespace, error)
}

type namespacesService struct {
	client *client
}

func (n *namespacesService) GetNamespace(ctx context.Context, namespaceName string) (*platformapi.Namespace, error) {
	namespaceRes, res, err := n.client.apiClient.NamespacesApi.GetNamespace(ctx, namespaceName)
	if vvperrors.IsResponseError(res) {
		return nil, vvperrors.FormatResponseError(res)
	}

	if err != nil {
		return nil, err
	}

	return namespaceRes.Namespace, nil
}

func (n *namespacesService) CreateNamespace(ctx context.Context, namespace platformapi.Namespace) (*platformapi.Namespace, error) {
	// name must be prefixed on creation
	if strings.HasPrefix(namespace.Name, "namespaces/") {
		namespace.Name = "namespaces/" + namespace.Name
	}

	namespaceRes, res, err := n.client.apiClient.NamespacesApi.CreateNamespace(ctx, namespace)
	if vvperrors.IsResponseError(res) {
		return nil, vvperrors.FormatResponseError(res)
	}

	if err != nil {
		return nil, err
	}

	return namespaceRes.Namespace, nil
}

func (n *namespacesService) UpdateNamespace(ctx context.Context, namespaceName string, namespace platformapi.Namespace) (*platformapi.Namespace, error) {
	// name must be prefixed on update
	if strings.HasPrefix(namespace.Name, "namespaces/") {
		namespace.Name = "namespaces/" + namespace.Name
	}

	namespaceRes, res, err := n.client.apiClient.NamespacesApi.UpdateNamespace(ctx, namespace, namespaceName)
	if vvperrors.IsResponseError(res) {
		return nil, vvperrors.FormatResponseError(res)
	}

	if err != nil {
		return nil, err
	}

	return namespaceRes.Namespace, nil
}

func (n *namespacesService) DeleteNamespace(ctx context.Context, namespaceName string) (*platformapi.Namespace, error) {
	namespaceRes, res, err := n.client.apiClient.NamespacesApi.DeleteNamespace(ctx, namespaceName)
	if vvperrors.IsResponseError(res) {
		return nil, vvperrors.FormatResponseError(res)
	}

	if err != nil {
		return nil, err
	}

	return namespaceRes.Namespace, nil
}

// API Tokens

var tokenNamePattern = regexp.MustCompile("namespaces/.+/apitokens/.+")

func formatTokenName(namespaceName, name string) string {
	if tokenNamePattern.Match([]byte(name)) {
		return name
	}

	return fmt.Sprintf("namespaces/%s/apitokens/%s", namespaceName, name)
}


type ApiTokensService interface {
	GetApiToken(ctx context.Context, namespaceName, name string) (*platformapi.ApiToken, error)
	CreateApiToken(ctx context.Context, namespaceName string, token platformapi.ApiToken) (*platformapi.ApiToken, error)
	DeleteApiToken(ctx context.Context, namespaceName, name string) error
}

type apiTokensService struct {
	client *client
}

func (s *apiTokensService) GetApiToken(ctx context.Context, namespaceName, name string) (*platformapi.ApiToken, error) {
	tokenRes, res, err := s.client.apiClient.ApiTokensApi.GetApiToken(ctx, name, namespaceName)
	if vvperrors.IsResponseError(res) {
		return nil, vvperrors.FormatResponseError(res)
	}

	if err != nil {
		return nil, err
	}

	return tokenRes.ApiToken, nil
}

func (s *apiTokensService) CreateApiToken(ctx context.Context, namespaceName string, apiToken platformapi.ApiToken) (*platformapi.ApiToken, error) {
	// name must be prefixed on creation
	apiToken.Name = formatTokenName(namespaceName, apiToken.Name)

	tokenRes, res, err := s.client.apiClient.ApiTokensApi.CreateApiToken(ctx, apiToken, namespaceName)
	if vvperrors.IsResponseError(res) {
		return nil, vvperrors.FormatResponseError(res)
	}

	if err != nil {
		return nil, err
	}

	return tokenRes.ApiToken, nil
}

func (s *apiTokensService) DeleteApiToken(ctx context.Context, namespaceName, name string) error {
	_, res, err := s.client.apiClient.ApiTokensApi.DeleteApiToken(ctx, name, namespaceName)
	if vvperrors.IsResponseError(res) {
		return vvperrors.FormatResponseError(res)
	}

	if err != nil {
		return err
	}

	return nil
}
