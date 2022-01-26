/*
 * Ververica Platform API
 *
 * The Ververica Platform APIs, excluding Application Manager.
 *
 * API version: 2.4.3
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package platformapi

type JobGraphSpec struct {
	AllowNonRestoredState  bool              `json:"allowNonRestoredState,omitempty"`
	FlinkVersion           string            `json:"flinkVersion,omitempty"`
	FullFlinkConfiguration map[string]string `json:"fullFlinkConfiguration,omitempty"`
	JobId                  string            `json:"jobId,omitempty"`
	SavepointLocation      string            `json:"savepointLocation,omitempty"`
	SqlStatement           string            `json:"sqlStatement,omitempty"`
	UserFlinkConfiguration map[string]string `json:"userFlinkConfiguration,omitempty"`
}
