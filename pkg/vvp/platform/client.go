package platform

import (
	"context"
	"fmt"
	vvperrors "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/errors"
	platformapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/platform-api"
	"regexp"
	"strings"
)

type Client interface {
	Namespaces() NamespacesService
	APITokens() APITokensService
}

type client struct {
	apiClient         *platformapi.APIClient
	namespacesService NamespacesService
	apiTokensService  APITokensService
}

func NewClient(config *platformapi.Configuration) Client {
	return &client{
		apiClient: platformapi.NewAPIClient(config),
	}
}

func (c *client) Namespaces() NamespacesService {
	if c.namespacesService == nil {
		c.namespacesService = &namespacesService{client: c}
	}
	return c.namespacesService
}

func (c *client) APITokens() APITokensService {
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
	namespaceRes, res, err := n.client.apiClient.NamespacesApi.GetNamespaceUsingGET(ctx, namespaceName)
	if vvperrors.IsResponseError(res) {
		return nil, vvperrors.FormatResponseError(res, err)
	}

	if err != nil {
		return nil, err
	}

	return namespaceRes.Namespace, nil
}

func (n *namespacesService) CreateNamespace(ctx context.Context, namespace platformapi.Namespace) (*platformapi.Namespace, error) {
	// name must be prefixed on creation
	if !strings.HasPrefix(namespace.Name, "namespaces/") {
		namespace.Name = "namespaces/" + namespace.Name
	}

	namespaceRes, res, err := n.client.apiClient.NamespacesApi.CreateNamespaceUsingPOST(ctx, namespace)
	if vvperrors.IsResponseError(res) {
		return nil, vvperrors.FormatResponseError(res, err)
	}

	if err != nil {
		return nil, err
	}

	return namespaceRes.Namespace, nil
}

func (n *namespacesService) UpdateNamespace(ctx context.Context, namespaceName string, namespace platformapi.Namespace) (*platformapi.Namespace, error) {
	// name must be prefixed on update
	if !strings.HasPrefix(namespace.Name, "namespaces/") {
		namespace.Name = "namespaces/" + namespace.Name
	}

	namespaceRes, res, err := n.client.apiClient.NamespacesApi.UpdateNamespaceUsingPUT(ctx, namespace, namespaceName)
	if vvperrors.IsResponseError(res) {
		return nil, vvperrors.FormatResponseError(res, err)
	}

	if err != nil {
		return nil, err
	}

	return namespaceRes.Namespace, nil
}

func (n *namespacesService) DeleteNamespace(ctx context.Context, namespaceName string) (*platformapi.Namespace, error) {
	namespaceRes, res, err := n.client.apiClient.NamespacesApi.DeleteNamespaceUsingDELETE(ctx, namespaceName)
	if vvperrors.IsResponseError(res) {
		return nil, vvperrors.FormatResponseError(res, err)
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

type APITokensService interface {
	GetAPIToken(ctx context.Context, namespaceName, name string) (*platformapi.ApiToken, error)
	CreateAPIToken(ctx context.Context, namespaceName string, token platformapi.ApiToken) (*platformapi.ApiToken, error)
	DeleteAPIToken(ctx context.Context, namespaceName, name string) error
}

type apiTokensService struct {
	client *client
}

func (s *apiTokensService) GetAPIToken(ctx context.Context, namespaceName, name string) (*platformapi.ApiToken, error) {
	tokenRes, res, err := s.client.apiClient.ApiTokensApi.GetApiTokenUsingGET(ctx, name, namespaceName)
	if vvperrors.IsResponseError(res) {
		return nil, vvperrors.FormatResponseError(res, err)
	}

	if err != nil {
		return nil, err
	}

	return tokenRes.ApiToken, nil
}

func (s *apiTokensService) CreateAPIToken(ctx context.Context, namespaceName string, apiToken platformapi.ApiToken) (*platformapi.ApiToken, error) {
	// name must be prefixed on creation
	apiToken.Name = formatTokenName(namespaceName, apiToken.Name)

	tokenRes, res, err := s.client.apiClient.ApiTokensApi.CreateApiTokenUsingPOST(ctx, apiToken, namespaceName)
	if vvperrors.IsResponseError(res) {
		return nil, vvperrors.FormatResponseError(res, err)
	}

	if err != nil {
		return nil, err
	}

	return tokenRes.ApiToken, nil
}

func (s *apiTokensService) DeleteAPIToken(ctx context.Context, namespaceName, name string) error {
	_, res, err := s.client.apiClient.ApiTokensApi.DeleteApiTokenUsingDELETE(ctx, name, namespaceName)
	if vvperrors.IsResponseError(res) {
		return vvperrors.FormatResponseError(res, err)
	}

	if err != nil {
		return err
	}

	return nil
}
