package converters

import (
	"encoding/json"
	"errors"

	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"
)

// DeploymentMetadataToNative converts a Ververica Platform deployment into its native K8s representation
func DeploymentMetadataToNative(deploymentMetadata vpAPI.DeploymentMetadata) (v1beta1.VpMetadata, error) {
	var vpMetadata v1beta1.VpMetadata
	metadataJSON, err := json.Marshal(deploymentMetadata)
	if err != nil {
		return vpMetadata, errors.New("cannot encode Metadata: " + err.Error())
	}

	// now unmarshal it into the platform model
	if err = json.Unmarshal(metadataJSON, &vpMetadata); err != nil {
		return vpMetadata, errors.New("cannot encode VpDeployment Metadata: " + err.Error())
	}

	return vpMetadata, nil
}

// DeploymentMetadataFromNative converts a native K8s VpDeployment to the Ververica Platform's representation
func DeploymentMetadataFromNative(vpMetadata v1beta1.VpMetadata) (vpAPI.DeploymentMetadata, error) {
	var deploymentMetadata vpAPI.DeploymentMetadata
	vpMetadataJSON, err := json.Marshal(vpMetadata)
	if err != nil {
		return deploymentMetadata, errors.New("cannot encode VpDeployment Metadata: " + err.Error())
	}

	// now unmarshal it into the platform model
	if err = json.Unmarshal(vpMetadataJSON, &deploymentMetadata); err != nil {
		return deploymentMetadata, errors.New("cannot encode Deployment Metadata: " + err.Error())
	}

	return deploymentMetadata, nil
}
