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

// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
type V1AzureFileVolumeSource struct {
	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly bool `json:"readOnly,omitempty"`
	// the name of secret that contains Azure Storage Account Name and Key
	SecretName string `json:"secretName"`
	// Share Name
	ShareName string `json:"shareName"`
}