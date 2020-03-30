package nativeconverters

import (
	"encoding/json"
	"errors"
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	appmanagerapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"
)

// DeploymentRunningStatusToNative converts a
func DeploymentRunningStatusToNative(status *appmanagerapi.DeploymentStatusRunning) (*v1beta2.VpDeploymentRunningStatus, error) {
	if status == nil {
		return nil, nil
	}

	// slow, yes -- easy and flexible, also yes
	var vpStatus v1beta2.VpDeploymentRunningStatus
	statusJSON, err := json.Marshal(status)
	if err != nil {
		return &vpStatus, errors.New("cannot encode Running Status: " + err.Error())
	}

	// now unmarshal it into the platform model
	if err = json.Unmarshal(statusJSON, &vpStatus); err != nil {
		return &vpStatus, errors.New("cannot encode VpDeployment Running Status: " + err.Error())
	}

	return &vpStatus, nil
}

// DeploymentRunningStatusFromNative converts a native K8s VpDeployment to the Ververica Platform's representation
func DeploymentRunningStatusFromNative(vpStatus *v1beta2.VpDeploymentRunningStatus) (*appmanagerapi.DeploymentStatusRunning, error) {
	if vpStatus == nil {
		return nil, nil
	}

	var status appmanagerapi.DeploymentStatusRunning
	vpStatusJSON, err := json.Marshal(vpStatus)
	if err != nil {
		return &status, errors.New("cannot encode Running Status: " + err.Error())
	}

	// now unmarshal it into the platform model
	if err = json.Unmarshal(vpStatusJSON, &status); err != nil {
		return &status, errors.New("cannot encode VpDeployment Running Status: " + err.Error())
	}

	return &status, nil
}
