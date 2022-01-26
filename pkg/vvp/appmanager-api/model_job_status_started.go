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

type JobStatusStarted struct {
	FlinkJobId string `json:"flinkJobId,omitempty"`
	LastUpdateTime time.Time `json:"lastUpdateTime,omitempty"`
	ObservedFlinkJobRestarts int32 `json:"observedFlinkJobRestarts,omitempty"`
	ObservedFlinkJobStatus string `json:"observedFlinkJobStatus,omitempty"`
	StartedAt time.Time `json:"startedAt,omitempty"`
}
