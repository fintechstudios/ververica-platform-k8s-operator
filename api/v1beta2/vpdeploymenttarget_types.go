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

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VpKubernetesTarget allows a user to configure k8s specific options
type VpKubernetesTarget struct {
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// VpDeploymentTargetSpec allows a users to set defaults for deployments and configure K8s
type VpDeploymentTargetSpec struct {
	Kubernetes VpKubernetesTarget `json:"kubernetes"`
}

// VpDeploymentTargetObjectSpec defines the desired state of VpDeploymentTarget
type VpDeploymentTargetObjectSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	Metadata VpMetadata `json:"metadata,omitempty"`
	// +optional
	Spec VpDeploymentTargetSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Namespace",type="string",JSONPath=".spec.metadata.namespace"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// VpDeploymentTarget is the Schema for the vpdeploymenttargets API
type VpDeploymentTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec VpDeploymentTargetObjectSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// VpDeploymentTargetList contains a list of VpDeploymentTarget
type VpDeploymentTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpDeploymentTarget `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VpDeploymentTarget{}, &VpDeploymentTargetList{})
}
