package converters

import (
	"errors"

	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
)

func SavepointOriginToNative(savepointOrigin string) (v1beta1.SavepointOrigin, error) {
	switch savepointOrigin {
	case string(v1beta1.UserRequestOrigin):
		return v1beta1.UserRequestOrigin, nil
	case string(v1beta1.SuspendAndUpgradeOrigin):
		return v1beta1.SuspendAndUpgradeOrigin, nil
	case string(v1beta1.SuspendOrigin):
		return v1beta1.SuspendOrigin, nil
	case string(v1beta1.CopiedOrigin):
		return v1beta1.CopiedOrigin, nil
	default:
		return "", errors.New("origin must be one of: USER_REQUEST, SUSPEND_AND_UPGRADE, SUSPEND, COPIED")
	}
}

func SavepointOriginFromNative(vpSavepointOrigin v1beta1.SavepointOrigin) (string, error) {
	switch vpSavepointOrigin {
	case v1beta1.UserRequestOrigin:
		return string(v1beta1.UserRequestOrigin), nil
	case v1beta1.SuspendAndUpgradeOrigin:
		return string(v1beta1.SuspendAndUpgradeOrigin), nil
	case v1beta1.SuspendOrigin:
		return string(v1beta1.SuspendOrigin), nil
	case v1beta1.CopiedOrigin:
		return string(v1beta1.CopiedOrigin), nil
	default:
		return "", errors.New("origin must be one of: USER_REQUEST, SUSPEND_AND_UPGRADE, SUSPEND, COPIED")
	}
}

func SavepointStateToNative(state string) (v1beta1.SavepointState, error) {
	switch state {
	case string(v1beta1.StartedSavepointState):
		return v1beta1.StartedSavepointState, nil
	case string(v1beta1.CompletedSavepointState):
		return v1beta1.CompletedSavepointState, nil
	case string(v1beta1.FailedSavepointState):
		return v1beta1.FailedSavepointState, nil
	default:
		return "", errors.New("state must be one of: STARTED, COMPLETED, FAILED")
	}
}

func SavepointStateFromNative(state v1beta1.SavepointState) (string, error) {
	switch state {
	case v1beta1.StartedSavepointState:
		return string(v1beta1.StartedSavepointState), nil
	case v1beta1.CompletedSavepointState:
		return string(v1beta1.CompletedSavepointState), nil
	case v1beta1.FailedSavepointState:
		return string(v1beta1.FailedSavepointState), nil
	default:
		return "", errors.New("state must be one of: STARTED, COMPLETED, FAILED")
	}
}
