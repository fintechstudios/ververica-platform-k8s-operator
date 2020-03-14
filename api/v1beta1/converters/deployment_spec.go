package converters

import (
	"encoding/json"
	"errors"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-operator/internal/vvp/appmanager-api-client"
)

// Only difference from Ververica Platform
// - spec.template.resources{}.cpu => Quantity in K8s, number in VP
// Need to remove these first, if they exist, or an unmarshalling error will occur

// DeploymentSpecToNative converts a Ververica Platform deployment into its native K8s representation
func DeploymentSpecToNative(deploymentSpec vpAPI.DeploymentSpec) (ververicaplatformv1beta1.VpDeploymentSpec, error) {
	var vpResources map[string]ververicaplatformv1beta1.VpResourceSpec
	if deploymentSpec.Template == nil ||
		deploymentSpec.Template.Spec == nil {
		return ververicaplatformv1beta1.VpDeploymentSpec{}, errors.New("invalid deployment spec: must provide template")
	}

	if deploymentSpec.Template.Spec.Resources != nil {
		vpResources, _ = ResourcesToNative(deploymentSpec.Template.Spec.Resources)
		deploymentSpec.Template.Spec.Resources = nil // don't try to marshal it
	}

	var vpDeploymentSpec ververicaplatformv1beta1.VpDeploymentSpec
	specJSON, err := json.Marshal(deploymentSpec)
	if err != nil {
		return vpDeploymentSpec, errors.New("cannot encode VpDeployment spec: " + err.Error())
	}

	// now unmarshal it into the platform model
	if err = json.Unmarshal(specJSON, &vpDeploymentSpec); err != nil {
		return vpDeploymentSpec, errors.New("cannot encode VpDeployment spec: " + err.Error())
	}

	if vpResources != nil {
		vpDeploymentSpec.Template.Spec.Resources = vpResources
	}

	return vpDeploymentSpec, nil
}

// DeploymentSpecFromNative converts a native K8s VpDeployment to the Ververica Platform's representation
func DeploymentSpecFromNative(vpDeploymentSpec ververicaplatformv1beta1.VpDeploymentSpec) (vpAPI.DeploymentSpec, error) {
	var resources map[string]vpAPI.ResourceSpec
	if vpDeploymentSpec.Template == nil ||
		vpDeploymentSpec.Template.Spec == nil {
		return vpAPI.DeploymentSpec{}, errors.New("invalid deployment spec: must provide template")
	}

	if vpDeploymentSpec.Template.Spec.Resources != nil {
		// Replace the resources with the corrected
		resources, _ = ResourcesFromNative(vpDeploymentSpec.Template.Spec.Resources)
		// remove so it is not marshalled
		vpDeploymentSpec.Template.Spec.Resources = nil
	}

	var deploymentSpec vpAPI.DeploymentSpec
	vpSpecJSON, err := json.Marshal(vpDeploymentSpec)
	if err != nil {
		return deploymentSpec, errors.New("cannot encode VpDeployment spec: " + err.Error())
	}

	// now unmarshal it into the platform model
	if err = json.Unmarshal(vpSpecJSON, &deploymentSpec); err != nil {
		return deploymentSpec, errors.New("cannot encode Deployment spec: " + err.Error())
	}

	if resources != nil {
		deploymentSpec.Template.Spec.Resources = resources
	}

	return deploymentSpec, nil
}
