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

type LibraStatus struct {
	LastActionTime time.Time `json:"lastActionTime,omitempty"`
	Message        string    `json:"message,omitempty"`
	Metrics        string    `json:"metrics,omitempty"`
	UpdateTime     time.Time `json:"updateTime,omitempty"`
}
