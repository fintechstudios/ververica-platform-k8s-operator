package nativeconverters

import (
	"errors"

	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	platformapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/platform-api"
)

var ErrorInvalidNamespaceLifecyclePhase = errors.New("origin must be one of: LIFECYCLE_PHASE_ACTIVE, LIFECYCLE_PHASE_TERMINATING, UNRECOGNIZED, LIFECYCLE_PHASE_INVALID")

func NamespaceLifecyclePhaseToNative(phase string) (v1beta1.NamespaceLifecyclePhase, error) {
	switch phase {
	case string(v1beta1.InvalidNamespaceLifecyclePhase):
		return v1beta1.InvalidNamespaceLifecyclePhase, nil
	case string(v1beta1.ActiveNamespaceLifecyclePhase):
		return v1beta1.ActiveNamespaceLifecyclePhase, nil
	case string(v1beta1.TerminatingNamespaceLifecyclePhase):
		return v1beta1.TerminatingNamespaceLifecyclePhase, nil
	case string(v1beta1.UnrecognizedNamespaceLifecyclePhase):
		return v1beta1.UnrecognizedNamespaceLifecyclePhase, nil
	default:
		return "", ErrorInvalidNamespaceLifecyclePhase
	}
}

func NamespaceLifecyclePhaseFromNative(vpPhase v1beta1.NamespaceLifecyclePhase) (string, error) {
	switch vpPhase {
	case v1beta1.InvalidNamespaceLifecyclePhase:
		return string(v1beta1.InvalidNamespaceLifecyclePhase), nil
	case v1beta1.ActiveNamespaceLifecyclePhase:
		return string(v1beta1.ActiveNamespaceLifecyclePhase), nil
	case v1beta1.TerminatingNamespaceLifecyclePhase:
		return string(v1beta1.TerminatingNamespaceLifecyclePhase), nil
	case v1beta1.UnrecognizedNamespaceLifecyclePhase:
		return string(v1beta1.UnrecognizedNamespaceLifecyclePhase), nil
	default:
		return "", ErrorInvalidNamespaceLifecyclePhase
	}
}

func NamespaceRoleBindingsFromNative(nativeBindings []v1beta1.NamespaceRoleBinding) []platformapi.RoleBinding {
	bindings := make([]platformapi.RoleBinding, len(nativeBindings))
	for i, nativeBinding := range nativeBindings {
		bindings[i] = platformapi.RoleBinding{
			Members: nativeBinding.Members,
			Role:    nativeBinding.Role,
		}
	}
	return bindings
}

func NamespaceRoleBindingsToNative(bindings []platformapi.RoleBinding) []v1beta1.NamespaceRoleBinding {
	nativeBindings := make([]v1beta1.NamespaceRoleBinding, len(bindings))
	for i, nativeBinding := range bindings {
		nativeBindings[i] = v1beta1.NamespaceRoleBinding{
			Members: nativeBinding.Members,
			Role:    nativeBinding.Role,
		}
	}
	return nativeBindings
}
