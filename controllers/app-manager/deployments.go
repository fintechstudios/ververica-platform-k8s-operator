package appmanager

import (
	"context"
	"errors"

	"github.com/fintechstudios/ververica-platform-k8s-operator/controllers/utils"

	appmanager "github.com/fintechstudios/ververica-platform-k8s-operator/appmanager-api-client"
)

// GetDeploymentByName fetches a deployment from the VP by namespace and name
func GetDeploymentByName(ctx context.Context, apiClient *appmanager.APIClient, namespace string, name string) (appmanager.Deployment, error) {
	var deployment appmanager.Deployment
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
