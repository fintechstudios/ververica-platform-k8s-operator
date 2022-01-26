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

type ResourceListOfJob struct {
	ApiVersion string `json:"apiVersion,omitempty"`
	Items []Job `json:"items,omitempty"`
	Kind string `json:"kind,omitempty"`
	Metadata *ResourceListMetadata `json:"metadata,omitempty"`
}
