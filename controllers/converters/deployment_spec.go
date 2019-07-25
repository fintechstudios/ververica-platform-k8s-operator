package converters

import (
	"encoding/json"
	"errors"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
)

// Only difference from Ververica Platform
// - spec.template.resources{}.cpu => Quantity in K8s, number in VP
// Need to remove these first, if they exist, or an unmarshalling error will occur

//DeploymentSpecToNative converts a Ververica Platform deployment into its native K8s representation
func DeploymentSpecToNative(deploymentSpec vpAPI.DeploymentSpec) (ververicaplatformv1beta1.VpDeploymentSpec, error) {
	var vpResources map[string]ververicaplatformv1beta1.VpResourceSpec
	if deploymentSpec.Template.Spec.Resources != nil {
		vpResources, _ = ResourcesToNative(deploymentSpec.Template.Spec.Resources)
		deploymentSpec.Template.Spec.Resources = nil
	}


	specJson, err := json.Marshal(deploymentSpec)
	// now unmarshal it into the platform model
	var vpDeploymentSpec ververicaplatformv1beta1.VpDeploymentSpec
	if err = json.Unmarshal(specJson, &vpDeploymentSpec); err != nil {
		return vpDeploymentSpec, errors.New("cannot encode VpDeployment spec: " + err.Error())
	}

	if vpResources != nil {
		vpDeploymentSpec.Template.Spec.Resources = vpResources
	}

	return vpDeploymentSpec, nil
}

// DeploymentSpecFromNative converts a native K8s VpDeployment to the Ververica Platform's representation
func DeploymentSpecFromNative(vpDeploymentSpec ververicaplatformv1beta1.VpDeploymentSpec) (vpAPI.DeploymentSpec, error) {
	vpTemplate := vpDeploymentSpec.Template
	var resources map[string]vpAPI.ResourceSpec
	if vpDeploymentSpec.Template != nil && vpTemplate.Spec.Resources != nil {
		// Replace the resources with the corrected
		resources, _ = ResourcesFromNative(vpTemplate.Spec.Resources)
		// remove so it is not marshalled
		vpDeploymentSpec.Template.Spec.Resources = nil
	}

	vpSpecJson, err := json.Marshal(vpDeploymentSpec)
	// now unmarshal it into the platform model
	var deploymentSpec vpAPI.DeploymentSpec
	if err = json.Unmarshal(vpSpecJson, &deploymentSpec); err != nil {
		return deploymentSpec, errors.New("cannot encode Deployment spec: " + err.Error())
	}

	if resources != nil {
		deploymentSpec.Template.Spec.Resources = resources
	}

	return deploymentSpec, nil
}
