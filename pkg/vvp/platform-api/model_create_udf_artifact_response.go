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

type CreateUdfArtifactResponse struct {
	CollidingClasses []UdfClass   `json:"collidingClasses,omitempty"`
	UdfArtifact      *UdfArtifact `json:"udfArtifact,omitempty"`
}
