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

import (
	"time"
)

type UdfArtifact struct {
	CreateTime    time.Time  `json:"createTime,omitempty"`
	JarUpdateTime time.Time  `json:"jarUpdateTime,omitempty"`
	JarUrl        string     `json:"jarUrl,omitempty"`
	Name          string     `json:"name,omitempty"`
	UdfClasses    []UdfClass `json:"udfClasses,omitempty"`
	UpdateTime    time.Time  `json:"updateTime,omitempty"`
}
