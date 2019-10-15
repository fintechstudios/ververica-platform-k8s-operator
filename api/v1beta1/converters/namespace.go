package converters

import (
	"errors"

	"github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
)

func NamespaceStateToNative(savepointOrigin string) (v1beta1.NamespaceState, error) {
	switch savepointOrigin {
	case string(v1beta1.ActiveNamespaceState):
		return v1beta1.ActiveNamespaceState, nil
	case string(v1beta1.MarkedForDeletionNamespaceState):
		return v1beta1.MarkedForDeletionNamespaceState, nil
	default:
		return "", errors.New("origin must be one of: USER_REQUEST, SUSPEND_AND_UPGRADE, SUSPEND, COPIED")
	}
}

func NamespaceStateFromNative(vpSavepointOrigin v1beta1.NamespaceState) (string, error) {
	switch vpSavepointOrigin {
	case v1beta1.ActiveNamespaceState:
		return string(v1beta1.ActiveNamespaceState), nil
	case v1beta1.MarkedForDeletionNamespaceState:
		return string(v1beta1.MarkedForDeletionNamespaceState), nil
	default:
		return "", errors.New("origin must be one of: USER_REQUEST, SUSPEND_AND_UPGRADE, SUSPEND, COPIED")
	}
}
