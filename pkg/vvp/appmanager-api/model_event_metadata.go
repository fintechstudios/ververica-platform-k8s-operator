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

type EventMetadata struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`
	DeploymentId string `json:"deploymentId,omitempty"`
	Id string `json:"id,omitempty"`
	JobId string `json:"jobId,omitempty"`
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
	ResourceVersion int32 `json:"resourceVersion,omitempty"`
	SessionClusterId string `json:"sessionClusterId,omitempty"`
}
