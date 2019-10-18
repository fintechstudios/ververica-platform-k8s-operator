package utils

import (
	"strings"

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
		errMsg := strings.TrimSpace(err.Error())
		return errMsg == "404 Not Found" || // AppManager API
			errMsg == "404" || // Platform API
			apierrs.IsNotFound(err) // K8s
	}
}
