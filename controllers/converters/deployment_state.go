package converters

import (
	"errors"
	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
)

//DeploymentStateToNative converts a Ververica Platform deployment into its native K8s representation
func DeploymentStateToNative(state string) (ververicaplatformv1beta1.DeploymentState, error) {
	switch state {
	case string(ververicaplatformv1beta1.CancelledState):
		return ververicaplatformv1beta1.CancelledState, nil
	case string(ververicaplatformv1beta1.RunningState):
		return ververicaplatformv1beta1.RunningState, nil
	case string(ververicaplatformv1beta1.TransitioningState):
		return ververicaplatformv1beta1.TransitioningState, nil
	case string(ververicaplatformv1beta1.SuspendedState):
		return ververicaplatformv1beta1.SuspendedState, nil
	case string(ververicaplatformv1beta1.FailedState):
		return ververicaplatformv1beta1.FailedState, nil
	default:
		return "", errors.New("state must be one of: CANCELLED, RUNNING, TRANSITIONING, SUSPENDED, FAILED")
	}
}

