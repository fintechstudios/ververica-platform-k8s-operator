package converters

import (
	"encoding/json"
	"errors"
	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
)

//DeploymentMetadataToNative converts a Ververica Platform deployment into its native K8s representation
func DeploymentMetadataToNative(deploymentMetadata vpAPI.DeploymentMetadata) (ververicaplatformv1beta1.VpDeploymentMetadata, error) {
	var vpDeploymentMetadata ververicaplatformv1beta1.VpDeploymentMetadata
	metadataJSON, err := json.Marshal(deploymentMetadata)
	if err != nil {
		return vpDeploymentMetadata, errors.New("cannot encode VpDeployment Metadata: " + err.Error())
	}

	// now unmarshal it into the platform model
	if err = json.Unmarshal(metadataJSON, &vpDeploymentMetadata); err != nil {
		return vpDeploymentMetadata, errors.New("cannot encode VpDeployment Metadata: " + err.Error())
	}

	return vpDeploymentMetadata, nil
}

// DeploymentMetadataFromNative converts a native K8s VpDeployment to the Ververica Platform's representation
func DeploymentMetadataFromNative(vpDeploymentMetadata ververicaplatformv1beta1.VpDeploymentMetadata) (vpAPI.DeploymentMetadata, error) {
	var deploymentMetadata vpAPI.DeploymentMetadata
	vpMetadataJSON, err := json.Marshal(vpDeploymentMetadata)
	if err != nil {
		return deploymentMetadata, errors.New("cannot encode VpDeployment Metadata: " + err.Error())
	}

	// now unmarshal it into the platform model
	if err = json.Unmarshal(vpMetadataJSON, &deploymentMetadata); err != nil {
		return deploymentMetadata, errors.New("cannot encode Deployment Metadata: " + err.Error())
	}

	// time.Time doesn't serialize correctly, so map over manually
	deploymentMetadata.CreatedAt = vpDeploymentMetadata.CreatedAt.Time
	deploymentMetadata.ModifiedAt = vpDeploymentMetadata.ModifiedAt.Time

	return deploymentMetadata, nil
}
