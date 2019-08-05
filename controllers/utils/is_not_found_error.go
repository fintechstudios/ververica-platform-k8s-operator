package utils

import (
	apierrs "k8s.io/apimachinery/pkg/api/errors"
)

// IsNotFoundError returns if an error is a 404
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	switch err.(type) {
	// internal
	case DeploymentNotFoundError:
		return true
	// external
	default:
		return err.Error() == "404 Not Found" || // Swagger
			apierrs.IsNotFound(err) // K8s
	}
}
