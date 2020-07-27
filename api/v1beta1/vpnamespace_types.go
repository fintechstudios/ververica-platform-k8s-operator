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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type NamespaceRoleBinding struct {
	// +optional
	Members []string `json:"members,omitempty"`
	// +optional
	Role string `json:"role,omitempty"`
}

// VpNamespaceSpec defines the desired state of VpNamespace
type VpNamespaceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Kind and ApiVersion are mapped automatically
	// Status is mapped to a subresource

	// +optional
	RoleBindings []NamespaceRoleBinding `json:"roleBindings,omitempty"`
}

// NamespaceLifecyclePhase is the enum of all possible namespace lifecycle phase
// Only one of the following states may be specified.
// +kubebuilder:validation:Enum=LIFECYCLE_PHASE_INVALID;LIFECYCLE_PHASE_ACTIVE;LIFECYCLE_PHASE_TERMINATING;UNRECOGNIZED
type NamespaceLifecyclePhase string

const (
	InvalidNamespaceLifecyclePhase      = NamespaceLifecyclePhase("LIFECYCLE_PHASE_INVALID")
	ActiveNamespaceLifecyclePhase       = NamespaceLifecyclePhase("LIFECYCLE_PHASE_ACTIVE")
	TerminatingNamespaceLifecyclePhase  = NamespaceLifecyclePhase("LIFECYCLE_PHASE_TERMINATING")
	UnrecognizedNamespaceLifecyclePhase = NamespaceLifecyclePhase("UNRECOGNIZED")
)

// VpNamespaceStatus defines the observed state of VpNamespace
type VpNamespaceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	LifecyclePhase NamespaceLifecyclePhase `json:"lifecyclePhase,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="LifecyclePhase",type="string",JSONPath=".status.lifecyclePhase"

// VpNamespace is the Schema for the vpnamespaces API
type VpNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpNamespaceSpec   `json:"spec,omitempty"`
	Status            VpNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpNamespaceList contains a list of VpNamespace
type VpNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpNamespace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VpNamespace{}, &VpNamespaceList{})
}
