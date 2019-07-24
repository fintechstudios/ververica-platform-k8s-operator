package utils

// DefaultStrIfEmpty returns a default value if the string is empty
func DefaultStrIfEmpty(str string, defaultVal string) string {
	if len(str) == 0 {
		return defaultVal
	}
	return str
}
