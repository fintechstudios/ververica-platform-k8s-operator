package nativeconverters

import (
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/annotations"
	appmanagerapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"
)

// DeploymentFromNative converts a native K8s VpDeployment to the Ververica Platform's representation
func DeploymentFromNative(vpDeployment v1beta2.VpDeployment) (appmanagerapi.Deployment, error) {
	deployment := appmanagerapi.Deployment{
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

	if deployment.Status, err = DeploymentStatusFromNative(vpDeployment.Status); err != nil {
		return deployment, err
	}

	if vpDeployment.Status != nil && len(vpDeployment.Status.State) > 0 {
		// we've got some state
		state, err := DeploymentStateFromNative(vpDeployment.Status.State)
		if err != nil {
			return deployment, err
		}

		runningStatus, err := DeploymentRunningStatusFromNative(vpDeployment.Status.Running)
		if err != nil {
			return deployment, err
		}

		deployment.Status = &appmanagerapi.DeploymentStatus{
			State:   state,
			Running: runningStatus,
		}
	}

	return deployment, nil
}
