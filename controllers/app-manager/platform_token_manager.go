package appmanager

import (
	"context"
	"fmt"

	platformApiClient "github.com/fintechstudios/ververica-platform-k8s-operator/platform-api-client"
)

func formatTokenNameForNamespace(name, namespace string) string {
	return fmt.Sprintf("namespaces/%s/apitokens/%s", namespace, name)
}

// PlatformTokenManager manages creation / deletion / querying of Platform API Tokens
type PlatformTokenManager struct {
	PlatformAPIClient *platformApiClient.APIClient
}

func (p *PlatformTokenManager) TokenExists(ctx context.Context, name, namespace string) (bool, error) {
	_, res, err := p.PlatformAPIClient.ApiTokensApi.GetApiToken(ctx, name, namespace)
	if res != nil && (res.StatusCode == 404 || res.StatusCode == 403) {
		return false, nil
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *PlatformTokenManager) CreateToken(ctx context.Context, name, role, namespace string) (string, error) {
	createRes, _, err := p.PlatformAPIClient.ApiTokensApi.CreateApiToken(ctx, platformApiClient.ApiToken{
		Name: formatTokenNameForNamespace(name, namespace),
		Role: role,
	}, namespace)

	if err != nil {
		return "", err
	}

	return createRes.ApiToken.Secret, nil
}

func (p *PlatformTokenManager) RemoveToken(ctx context.Context, name, namespace string) (bool, error) {
	_, res, err := p.PlatformAPIClient.ApiTokensApi.DeleteApiToken(ctx, name, namespace)
	if res != nil && res.StatusCode == 404 {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
