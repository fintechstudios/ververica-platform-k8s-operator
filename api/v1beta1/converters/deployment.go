package converters

import (
	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/appmanager-api-client"
	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/annotations"
)

// DeploymentFromNative converts a native K8s VpDeployment to the Ververica Platform's representation
func DeploymentFromNative(vpDeployment ververicaplatformv1beta1.VpDeployment) (vpAPI.Deployment, error) {
	deployment := vpAPI.Deployment{
		Kind:       "Deployment",
		ApiVersion: "v1",
	}

	deploymentSpec, err := DeploymentSpecFromNative(vpDeployment.Spec.Spec)
	if err != nil {
		return deployment, err
	}

	deploymentMeta, err := DeploymentMetadataFromNative(vpDeployment.Spec.Metadata)
	if err != nil {
		return deployment, err
	}
	deploymentMeta.Id = annotations.Get(vpDeployment.Annotations, annotations.ID)
	deploymentMeta.Name = vpDeployment.Name
	deployment.Metadata = &deploymentMeta

	if annotations.Has(vpDeployment.Annotations, annotations.DeploymentTargetID) {
		deploymentSpec.DeploymentTargetId = annotations.Get(vpDeployment.Annotations, annotations.DeploymentTargetID)
	}
	deployment.Spec = &deploymentSpec

	if len(vpDeployment.Status.State) > 0 {
		state, err := DeploymentStateFromNative(vpDeployment.Status.State)
		if err != nil {
			return deployment, err
		}

		deployment.Status = &vpAPI.DeploymentStatus{State: state}
	}

	return deployment, nil
}
