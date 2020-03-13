package converters

import (
	"errors"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-operator/internal/appmanager-api-client"
)

// DeploymentTargetPatchSetFromNative converts a deployment k8s patch set to a platform patch set
func DeploymentTargetPatchSetFromNative(vpPatchSet []ververicaplatformv1beta1.JSONPatchGeneric) ([]vpAPI.JsonPatchGeneric, error) {
	patchSet := make([]vpAPI.JsonPatchGeneric, len(vpPatchSet))
	for i, patch := range vpPatchSet {
		patchSet[i] = vpAPI.JsonPatchGeneric{
			Op:    patch.Op,
			Path:  patch.Path,
			From:  patch.From,
			Value: patch.Value,
		}
	}

	return patchSet, nil
}

func parseDeploymentPatchValue(value vpAPI.Any) (*string, bool) {
	if value == nil {
		return nil, true
	}

	switch val := value.(type) {
	case string:
		return &val, true
	default:
		return nil, false
	}
}

// DeploymentTargetPatchSetToNative converts a deployment platform patch set to a k8s patch set
func DeploymentTargetPatchSetToNative(patchSet []vpAPI.JsonPatchGeneric) ([]ververicaplatformv1beta1.JSONPatchGeneric, error) {
	vpPatchSet := make([]ververicaplatformv1beta1.JSONPatchGeneric, len(patchSet))
	for i, patch := range patchSet {
		value, ok := parseDeploymentPatchValue(patch.Value)
		if !ok {
			return nil, errors.New("invalid patch value: only strings are supported")
		}

		vpPatchSet[i] = ververicaplatformv1beta1.JSONPatchGeneric{
			Op:    patch.Op,
			Path:  patch.Path,
			From:  patch.From,
			Value: value,
		}
	}

	return vpPatchSet, nil
}
