package appmanager

import (
	"context"
)

const NoOpToken = "no-op-token"

// NoOpTokenManager is a singleton instance of a do-nothing token manager.
var NoOpTokenManager = noopTokenManager{}

// noopTokenManager is a token manager that does nothing.
// Used for the Community Edition where API tokens are not supported.
type noopTokenManager struct {
}

// TokenExists always returns true
func (p *noopTokenManager) TokenExists(ctx context.Context, namespaceName, name string) (bool, error) {
	return true, nil
}

// CreateToken always returns the constant NoOpToken value
func (p *noopTokenManager) CreateToken(ctx context.Context, namespaceName, name, role string) (string, error) {
	return NoOpToken, nil
}

// RemoveToken always returns true
func (p *noopTokenManager) RemoveToken(ctx context.Context, namespaceName, name string) (bool, error) {
	return true, nil
}
