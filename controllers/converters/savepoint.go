package converters

import (
	"encoding/json"
	"errors"

	"github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/appmanager-api-client"
)

func SavepointSpecToNative(savepointSpec vpAPI.SavepointSpec) (v1beta1.VpSavepointSpec, error) {
	return v1beta1.VpSavepointSpec{
		SavepointLocation: savepointSpec.SavepointLocation,
		FlinkSavepointID:  savepointSpec.FlinkSavepointId,
	}, nil
}

func SavepointSpecFromNative(vpSavepointSpec v1beta1.VpSavepointSpec) (vpAPI.SavepointSpec, error) {
	return vpAPI.SavepointSpec{
		SavepointLocation: vpSavepointSpec.SavepointLocation,
		FlinkSavepointId:  vpSavepointSpec.FlinkSavepointID,
	}, nil
}

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

func SavepointMetadataToNative(savepointMeta vpAPI.SavepointMetadata) (v1beta1.VpSavepointMetadata, error) {
	var vpSavepointMeta v1beta1.VpSavepointMetadata
	metadataJSON, err := json.Marshal(savepointMeta)
	if err != nil {
		return vpSavepointMeta, errors.New("cannot encode Metadata: " + err.Error())
	}

	// now unmarshal it into the platform model
	if err = json.Unmarshal(metadataJSON, &vpSavepointMeta); err != nil {
		return vpSavepointMeta, errors.New("cannot encode VpSavepoint Metadata: " + err.Error())
	}

	origin, err := SavepointOriginToNative(savepointMeta.Origin)
	if err != nil {
		return vpSavepointMeta, err
	}
	vpSavepointMeta.Origin = origin

	return vpSavepointMeta, nil
}

func SavepointMetadataFromNative(vpSavepointMeta v1beta1.VpSavepointMetadata) (vpAPI.SavepointMetadata, error) {
	var savepointMeta vpAPI.SavepointMetadata

	vpMetadataJSON, err := json.Marshal(vpSavepointMeta)
	if err != nil {
		return savepointMeta, errors.New("cannot encode VpSavepoint Metadata: " + err.Error())
	}

	// now unmarshal it into the platform model
	if err = json.Unmarshal(vpMetadataJSON, &savepointMeta); err != nil {
		return savepointMeta, errors.New("cannot encode Savepoint Metadata: " + err.Error())
	}

	return savepointMeta, nil
}
