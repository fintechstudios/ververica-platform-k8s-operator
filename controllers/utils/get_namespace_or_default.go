package utils


var defaultNamespace = "default"

func GetNamespaceOrDefault(namespace *string) string {
	if namespace == nil || len(*namespace) == 0 {
		return defaultNamespace
	}
	return *namespace
}