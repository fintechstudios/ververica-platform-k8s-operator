package utils

import (
	appManagerApiClient "github.com/fintechstudios/ververica-platform-k8s-operator/appmanager-api-client"
	platformApiClient "github.com/fintechstudios/ververica-platform-k8s-operator/platform-api-client"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// IsNotFoundError returns if an error is a 404
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	switch err := err.(type) {
	// internal
	case DeploymentNotFoundError:
		return true
	case platformApiClient.GenericSwaggerError:
		return err.StatusCode() == 404
	case appManagerApiClient.GenericSwaggerError:
		return err.StatusCode() == 404
	// external
	default:
		return apierrors.IsNotFound(err) // K8s
	}
}
