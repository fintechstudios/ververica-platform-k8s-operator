package platform

import (
	"context"
	"errors"
	vvperrors "github.com/fintechstudios/ververica-platform-k8s-operator/internal/vvp/errors"

	platformapi "github.com/fintechstudios/ververica-platform-k8s-operator/internal/vvp/platform-api-client"
)

// TokenManager manages creation / deletion / querying of Platform API Tokens
type TokenManager struct {
	PlatformClient Client
}

func (p *TokenManager) TokenExists(ctx context.Context, name, namespace string) (bool, error) {
	_, err := p.PlatformClient.ApiTokens().GetApiToken(ctx, name, namespace)

	if errors.Is(err, vvperrors.ErrNotFound) || errors.Is(err, vvperrors.ErrForbidden) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *TokenManager) CreateToken(ctx context.Context, name, role, namespace string) (string, error) {
	token, err := p.PlatformClient.ApiTokens().CreateApiToken(ctx, namespace, platformapi.ApiToken{
		Name: name,
		Role: role,
	})

	if err != nil {
		return "", err
	}

	return token.Secret, nil
}

func (p *TokenManager) RemoveToken(ctx context.Context, name, namespace string) (bool, error) {
	err := p.PlatformClient.ApiTokens().DeleteApiToken(ctx, name, namespace)

	// We're ok if it's not found
	if errors.Is(err, vvperrors.ErrNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
