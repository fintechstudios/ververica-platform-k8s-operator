package nativeconverters

import (
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"
)

// DeploymentStatusFromNative converts a native K8s VpDeployment.Status to the Ververica Platform's representation
func DeploymentStatusFromNative(vpStatus *v1beta2.VpDeploymentStatus) (status *appmanagerapi.DeploymentStatus, err error) {
	if vpStatus == nil {
		return nil, nil
	}

	status = &appmanagerapi.DeploymentStatus{}

	if len(vpStatus.State) > 0 {
		if status.State, err = DeploymentStateFromNative(vpStatus.State); err != nil {
			return status, err
		}
	}

	if status.Running, err = DeploymentRunningStatusFromNative(vpStatus.Running); err != nil {
		return status, err
	}

	return status, nil
}

// DeploymentStatusToNative converts to a native K8s VpDeployment.Status from the Ververica Platform's representation
func DeploymentStatusToNative(status *appmanagerapi.DeploymentStatus) (vpStatus *v1beta2.VpDeploymentStatus, err error) {
	if status == nil {
		return nil, nil
	}

	vpStatus = &v1beta2.VpDeploymentStatus{}

	if len(status.State) > 0 {
		if vpStatus.State, err = DeploymentStateToNative(status.State); err != nil {
			return vpStatus, err
		}
	}

	if vpStatus.Running, err = DeploymentRunningStatusToNative(status.Running); err != nil {
		return vpStatus, err
	}

	return vpStatus, nil
}
