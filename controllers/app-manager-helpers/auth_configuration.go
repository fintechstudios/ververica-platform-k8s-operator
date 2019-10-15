package appManagerHelpers


type AuthConfiguration struct {

}

func NewAuthConfiguration() *AuthConfiguration {
	return &AuthConfiguration{}
}

func (authConf *AuthConfiguration) getTokenForNamespace(namespace string) {
	// search chain:
	// - environment in form VP_API_TOKEN_{NAMESPACE}
	// - k8s secret?
}