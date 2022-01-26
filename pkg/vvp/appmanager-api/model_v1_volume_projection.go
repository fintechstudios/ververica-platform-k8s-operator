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

// Projection that may be projected along with other supported volume types
type V1VolumeProjection struct {
	ConfigMap           *V1ConfigMapProjection           `json:"configMap,omitempty"`
	DownwardAPI         *V1DownwardApiProjection         `json:"downwardAPI,omitempty"`
	Secret              *V1SecretProjection              `json:"secret,omitempty"`
	ServiceAccountToken *V1ServiceAccountTokenProjection `json:"serviceAccountToken,omitempty"`
}
