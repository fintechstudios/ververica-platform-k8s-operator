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

type DeploymentCondition struct {
	LastTransitionTime time.Time `json:"lastTransitionTime,omitempty"`
	LastUpdateTime     time.Time `json:"lastUpdateTime,omitempty"`
	Message            string    `json:"message,omitempty"`
	Reason             string    `json:"reason,omitempty"`
	Status             string    `json:"status,omitempty"`
	Type_              string    `json:"type,omitempty"`
}
