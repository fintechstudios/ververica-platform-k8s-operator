package utils

// IgnoreNotFoundError returns nil for Swagger API 404 errors, otherwise the original error
func IgnoreNotFoundError(err error) error {
	if IsNotFoundError(err) {
		return nil
	}
	return err
}
