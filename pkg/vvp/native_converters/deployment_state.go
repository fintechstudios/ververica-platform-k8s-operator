package nativeconverters

import (
	"errors"
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
)

var ErrInvalidDeploymentState = errors.New("state must be one of: CANCELLED, RUNNING, TRANSITIONING, SUSPENDED, FINISHED, FAILED")

// DeploymentStateToNative converts a Ververica Platform deployment into its native K8s representation
func DeploymentStateToNative(state string) (v1beta2.VpDeploymentState, error) {
	switch state {
	case string(v1beta2.CancelledState):
		return v1beta2.CancelledState, nil
	case string(v1beta2.RunningState):
		return v1beta2.RunningState, nil
	case string(v1beta2.TransitioningState):
		return v1beta2.TransitioningState, nil
	case string(v1beta2.SuspendedState):
		return v1beta2.SuspendedState, nil
	case string(v1beta2.FailedState):
		return v1beta2.FailedState, nil
	case string(v1beta2.FinishedState):
		return v1beta2.FinishedState, nil
	default:
		return "", ErrInvalidDeploymentState
	}
}

// DeploymentStateFromNative converts a native K8s VpDeployment to the Ververica Platform's representation
func DeploymentStateFromNative(vpState v1beta2.VpDeploymentState) (string, error) {
	switch vpState {
	case v1beta2.CancelledState:
		return string(v1beta2.CancelledState), nil
	case v1beta2.RunningState:
		return string(v1beta2.RunningState), nil
	case v1beta2.TransitioningState:
		return string(v1beta2.TransitioningState), nil
	case v1beta2.SuspendedState:
		return string(v1beta2.SuspendedState), nil
	case v1beta2.FinishedState:
		return string(v1beta2.FinishedState), nil
	case v1beta2.FailedState:
		return string(v1beta2.FailedState), nil
	default:
		return "", ErrInvalidDeploymentState
	}
}
