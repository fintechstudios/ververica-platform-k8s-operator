package vpApiHelpers

import (
	"context"
	"errors"

	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/utils"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
)

// GetDeploymentByName
func GetDeploymentByName(apiClient *vpAPI.APIClient, ctx context.Context, namespace string, name string) (vpAPI.Deployment, error) {
	var deployment vpAPI.Deployment
	if len(namespace) == 0 || len(name) == 0 {
		return deployment, errors.New("namespace and name must not be empty")
	}

	deploymentsList, _, err := apiClient.DeploymentsApi.GetDeployments(ctx, namespace, nil)

	if err != nil {
		return deployment, err
	}

	for _, deployment = range deploymentsList.Items {
		if deployment.Metadata.Name == name {
			return deployment, nil
		}
	}

	return deployment, utils.DeploymentNotFoundError{Namespace: namespace, Name: name}
}
