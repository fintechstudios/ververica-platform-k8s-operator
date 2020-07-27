/*
Copyright 2020 FinTech Studios, Inc.

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
	SavepointLocation string `json:"savepointLocation"`
	FlinkSavepointID  string `json:"flinkSavepointId"`
}

// +kubebuilder:validation:Enum=USER_REQUEST;SUSPEND_AND_UPGRADE;SUSPEND;COPIED
type SavepointOrigin string

const (
	UserRequestOrigin       = SavepointOrigin("USER_REQUEST")
	SuspendAndUpgradeOrigin = SavepointOrigin("SUSPEND_AND_UPGRADE")
	SuspendOrigin           = SavepointOrigin("SUSPEND")
	CopiedOrigin            = SavepointOrigin("COPIED")
)

type VpSavepointMetadata struct {
	// +optional
	VpMetadata `json:",inline"`

	// Can be specified through the .spec.deploymentName
	// +optional
	DeploymentID string `json:"deploymentId,omitempty"`
}

// VpSavepointSpec defines the desired state of VpSavepoint
type VpSavepointObjectSpec struct {
	// +optional
	Metadata VpSavepointMetadata `json:"metadata,omitempty"`

	// +optional
	Spec VpSavepointSpec `json:"spec,omitempty"`

	// DeploymentName is an extension on the VP API
	// Must provide a spec.metadata.deploymentId if not set
	// +optional
	DeploymentName string `json:"deploymentName,omitempty"`
}

// +kubebuilder:validation:Enum=STARTED;COMPLETED;FAILED
type SavepointState string

const (
	StartedSavepointState   = SavepointState("STARTED")
	CompletedSavepointState = SavepointState("COMPLETED")
	FailedSavepointState    = SavepointState("FAILED")
)

// VpSavepointStatus defines the observed state of VpSavepoint
type VpSavepointStatus struct {
	State SavepointState `json:"state"`
	// +optional
	Failure VpSavepointFailure `json:"failure,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"

// VpSavepoint is the Schema for the vpsavepoints API
type VpSavepoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpSavepointObjectSpec `json:"spec,omitempty"`
	Status VpSavepointStatus     `json:"status,omitempty"`
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
