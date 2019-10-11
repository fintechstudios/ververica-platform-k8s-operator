/*
Copyright 2019 FinTech Studios, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type VpSavepointFailure struct {
	// +optional
	FailedAt string `json:"failedAt,omitempty"`
	// +optional
	Message string `json:"message,omitempty"`
	// +optional
	Reason string `json:"reason,omitempty"`
}

type VpSavepointSpec struct {
	// +optional
	SavepointLocation string `json:"savepointLocation,omitempty"`
	FlinkSavepointID  string `json:"flinkSavepointId"`
}

// +kubebuilder:validation:Enum=USER_REQUEST;SUSPEND_AND_UPGRADE;SUSPEND;COPIED
type savepointOrigin string

const (
	UserRequestOrigin       = savepointOrigin("USER_REQUEST")
	SuspendAndUpgradeOrigin = savepointOrigin("SUSPEND_AND_UPGRADE")
	SuspendOrigin           = savepointOrigin("SUSPEND")
	CopiedOrigin            = savepointOrigin("COPIED")
)

type VpSavepointMetadata struct {
	// +optional
	VpMetadata `json:",inline"`

	// Can be specified through the VpSavepointObjectSpec
	// +optional
	DeploymentID string `json:"deploymentId,omitempty"`

	// +optional
	JobID string `json:"jobId,omitempty"`

	// +optional
	Origin savepointOrigin `json:"origin,omitempty"`
}

// VpSavepointSpec defines the desired state of VpSavepoint
type VpSavepointObjectSpec struct {
	// +optional
	Metadata VpMetadata `json:"metadata,omitempty"`

	// +optional
	Spec VpSavepointSpec `json:"spec,omitempty"`

	// DeploymentName is an extension on the VP API
	// Must provide a spec.metadata.deploymentId if not set
	// +optional
	DeploymentName string `json:"deploymentName,omitempty"`
}

// +kubebuilder:validation:Enum=STARTED;COMPLETED;FAILED
type savepointState string

const (
	StartedSavepointState   = savepointState("STARTED")
	CompletedSavepointState = savepointState("COMPLETED")
	FailedSavepointState    = savepointState("FAILED")
)

// VpSavepointStatus defines the observed state of VpSavepoint
type VpSavepointStatus struct {
	State savepointState `json:"state"`
	// +optional
	Failure VpSavepointFailure `json:"failure,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VpSavepoint is the Schema for the vpsavepoints API
type VpSavepoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpSavepointSpec   `json:"spec,omitempty"`
	Status VpSavepointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpSavepointList contains a list of VpSavepoint
type VpSavepointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpSavepoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VpSavepoint{}, &VpSavepointList{})
}
