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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VpNamespaceMetadata represents all metadata from the VP API
type VpNamespaceMetadata struct {
	// +optional
	Name string `json:"name"`
	// +optional
	Id string `json:"id,omitempty"`
	// +optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	// +optional
	ModifiedAt *metav1.Time `json:"modifiedAt,omitempty"`
	// +optional
	ResourceVersion int32 `json:"resourceVersion,omitempty"`
}

// VpNamespaceSpec defines the desired state of VpNamespace
type VpNamespaceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Kind and ApiVersion are mapped automatically
	// Status is mapped to a subresource
	// +optional
	Metadata VpNamespaceMetadata `json:"metadata,omitempty"`
}

// VpNamespaceStatus defines the observed state of VpNamespace
type VpNamespaceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	State string `json:"state,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Id",type="string",JSONPath=".spec.metadata.id"
// +kubebuilder:printcolumn:name="ResourceVersion",type="integer",JSONPath=".spec.metadata.resourceVersion"
// +kubebuilder:printcolumn:name="Created",type="date",JSONPath=".spec.metadata.createdAt"
// +kubebuilder:printcolumn:name="Modified",type="date",JSONPath=".spec.metadata.modifiedAt"

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
