package utils

import apierrs "k8s.io/apimachinery/pkg/api/errors"

// IsNotFoundError returns if an error is a 404
func IsNotFoundError(err error) bool {
	return err != nil && (
		err.Error() == "404 Not Found" || // Swagger
		apierrs.IsNotFound(err)) // K8s
}
