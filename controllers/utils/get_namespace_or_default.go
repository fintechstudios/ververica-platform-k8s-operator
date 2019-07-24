package utils


const defaultNamespace = "default"

// GetNamespaceOrDefault will return either the given namespace, if valid, or the default Ververica Platform namespace
func GetNamespaceOrDefault(namespace string) string {
	if len(namespace) == 0 {
		return defaultNamespace
	}
	return namespace
}