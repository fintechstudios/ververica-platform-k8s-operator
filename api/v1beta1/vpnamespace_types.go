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

// VPNamespaceMetadata represents all metadata from the VP API
type VPNamespaceMetadata struct {
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

// VPNamespaceSpec defines the desired state of VPNamespace
type VPNamespaceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Kind and ApiVersion are mapped automatically
	// Status is mapped to a subresource
	// +optional
	Metadata VPNamespaceMetadata `json:"metadata,omitempty"`
}

// VPNamespaceStatus defines the observed state of VPNamespace
type VPNamespaceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// A list of pointers to currently running jobs.
	// +optional
	State string `json:"state,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Id",type="string",JSONPath=".spec.metadata.id"
// +kubebuilder:printcolumn:name="ResourceVersion",type="integer",JSONPath=".spec.metadata.resourceVersion"
// +kubebuilder:printcolumn:name="Created",type="date",JSONPath=".spec.metadata.createdAt"
// +kubebuilder:printcolumn:name="Modified",type="date",JSONPath=".spec.metadata.modifiedAt"

// VPNamespace is the Schema for the vpnamespaces API
type VPNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec VPNamespaceSpec `json:"spec,omitempty"`
	Status VPNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNamespaceList contains a list of VPNamespace
type VPNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNamespace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VPNamespace{}, &VPNamespaceList{})
}
