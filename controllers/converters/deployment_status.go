package converters

import (
	"errors"
	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
)

//DeploymentMetadataToNative converts a Ververica Platform deployment into its native K8s representation
func DeploymentStatusToNative(status vpAPI.DeploymentStatus) (ververicaplatformv1beta1.VpDeploymentStatus, error) {
	vpStatus := ververicaplatformv1beta1.VpDeploymentStatus{}

	switch status.State {
	case string(ververicaplatformv1beta1.CancelledState):
		vpStatus.State = ververicaplatformv1beta1.CancelledState
	case string(ververicaplatformv1beta1.RunningState):
		vpStatus.State = ververicaplatformv1beta1.RunningState
	case string(ververicaplatformv1beta1.TransitioningState):
		vpStatus.State = ververicaplatformv1beta1.TransitioningState
	case string(ververicaplatformv1beta1.SuspendedState):
		vpStatus.State = ververicaplatformv1beta1.SuspendedState
	case string(ververicaplatformv1beta1.FailedState):
		vpStatus.State = ververicaplatformv1beta1.FailedState
	default:
		return vpStatus, errors.New("state must be one of: CANCELLED, RUNNING, TRANSITIONING, SUSPENDED, FAILED")
	}

	return vpStatus, nil
}

// DeploymentMetadataFromNative converts a native K8s VpDeployment to the Ververica Platform's representation
func DeploymentStatusFromNative(vpStatus ververicaplatformv1beta1.VpDeploymentStatus) (vpAPI.DeploymentStatus, error) {
	status := vpAPI.DeploymentStatus{}

	switch vpStatus.State {
	case ververicaplatformv1beta1.CancelledState:
		status.State = string(ververicaplatformv1beta1.CancelledState)
	case ververicaplatformv1beta1.RunningState:
		status.State = string(ververicaplatformv1beta1.RunningState)
	case ververicaplatformv1beta1.TransitioningState:
		status.State = string(ververicaplatformv1beta1.TransitioningState)
	case ververicaplatformv1beta1.SuspendedState:
		status.State = string(ververicaplatformv1beta1.SuspendedState)
	case ververicaplatformv1beta1.FailedState:
		status.State = string(ververicaplatformv1beta1.FailedState)
	default:
		return status, errors.New("state must be one of: CANCELLED, RUNNING, TRANSITIONING, SUSPENDED, FAILED")
	}

	return status, nil
}
