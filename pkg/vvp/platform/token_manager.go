package platform

import (
	"context"
	"errors"
	vvperrors "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/errors"

	platformapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/platform-api"
)

// TokenManager manages creation / deletion / querying of Platform API Tokens
type TokenManager struct {
	PlatformClient Client
}

func (p *TokenManager) TokenExists(ctx context.Context, namespaceName, name string) (bool, error) {
	_, err := p.PlatformClient.ApiTokens().GetApiToken(ctx, namespaceName, name)

	if errors.Is(err, vvperrors.ErrNotFound) || errors.Is(err, vvperrors.ErrForbidden) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *TokenManager) CreateToken(ctx context.Context, namespaceName, name, role string) (string, error) {
	token, err := p.PlatformClient.ApiTokens().CreateApiToken(ctx, namespaceName, platformapi.ApiToken{
		Name: name,
		Role: role,
	})

	if err != nil {
		return "", err
	}

	return token.Secret, nil
}

func (p *TokenManager) RemoveToken(ctx context.Context, namespaceName, name string) (bool, error) {
	err := p.PlatformClient.ApiTokens().DeleteApiToken(ctx, namespaceName, name)

	// We're ok if it's not found
	if errors.Is(err, vvperrors.ErrNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
