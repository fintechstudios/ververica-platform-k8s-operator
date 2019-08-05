package utils

import (
	"fmt"
)

type DeploymentNotFoundError struct {
	Name string
	Namespace string
}

func (err DeploymentNotFoundError) Error() string {
	return fmt.Sprintf("Deployment not found in the Ververica Platform by name %s in namespace %s", err.Name, err.Namespace)
}
