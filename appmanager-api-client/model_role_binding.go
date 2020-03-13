/*
 * AppManager API
 *
 * HTTP REST API to connect to the AppManager
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appmanagerapi

type RoleBinding struct {
	Kind string `json:"kind,omitempty"`
	ApiVersion string `json:"apiVersion,omitempty"`
	Metadata *RoleBindingMetadata `json:"metadata"`
	Spec *RoleBindingSpec `json:"spec"`
}
