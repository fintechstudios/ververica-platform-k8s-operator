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

type UpdateUdfArtifactResponse struct {
	CollidingClasses []UdfClass   `json:"collidingClasses,omitempty"`
	MissingClasses   []UdfClass   `json:"missingClasses,omitempty"`
	UdfArtifact      *UdfArtifact `json:"udfArtifact,omitempty"`
}
