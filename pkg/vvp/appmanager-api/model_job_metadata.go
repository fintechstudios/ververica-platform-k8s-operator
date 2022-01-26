/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.4.3
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appmanagerapi

import (
	"time"
)

type JobMetadata struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	DeploymentId string `json:"deploymentId,omitempty"`
	DeploymentName string `json:"deploymentName,omitempty"`
	Id string `json:"id,omitempty"`
	ModifiedAt time.Time `json:"modifiedAt,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	ResourceVersion int32 `json:"resourceVersion,omitempty"`
	SessionClusterName string `json:"sessionClusterName,omitempty"`
	TerminatedAt time.Time `json:"terminatedAt,omitempty"`
}
