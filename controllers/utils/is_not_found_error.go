package utils

// IsNotFoundError returns if Swagger API error is a 404
func IsNotFoundError(err error) bool {
	return err != nil && err.Error() == "404 Not Found"
}
