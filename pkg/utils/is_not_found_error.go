package utils

import (
	"errors"
	vvperrors "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// IsNotFoundError returns if an error is a 404
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, vvperrors.ErrNotFound) ||
		apierrors.IsNotFound(err) // k8s
}
