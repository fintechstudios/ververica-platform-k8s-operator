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

// JSONPatchGeneric is a generic JSON-Patch object
type JSONPatchGeneric struct {
	Op   string `json:"op"`
	Path string `json:"path"`
	// TODO: support any type of JSON as an interface
	// 		 https://github.com/kubernetes-sigs/kubebuilder/issues/528
	// +optional
	Value string `json:"value,omitempty"`
	// +optional
	From string `json:"from,omitempty"`
}

// VpDeploymentTargetMetadata represents all metadata from the VP API
type VpDeploymentTargetMetadata struct {
	// Taken from the K8s resource
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// +optional
	ID string `json:"id,omitempty"`
	// +optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	// +optional
	ModifiedAt *metav1.Time `json:"modifiedAt,omitempty"`
	// +optional
	ResourceVersion int32 `json:"resourceVersion,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
}

// VpKubernetesTarget allows a user to configure k8s specific options
type VpKubernetesTarget struct {
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// VpDeploymentTargetSpec allows a users to set defaults for deployments and configure K8s
type VpDeploymentTargetSpec struct {
	Kubernetes VpKubernetesTarget `json:"kubernetes"`
	// +optional
	DeploymentPatchSet []JSONPatchGeneric `json:"deploymentPatchSet,omitempty"`
}

// VpDeploymentTargetObjectSpec defines the desired state of VpDeploymentTarget
type VpDeploymentTargetObjectSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	Metadata VpDeploymentTargetMetadata `json:"metadata,omitempty"`
	// +optional
	Spec VpDeploymentTargetSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Id",type="string",JSONPath=".spec.metadata.id"
// +kubebuilder:printcolumn:name="Namespace",type="string",JSONPath=".spec.metadata.namespace"
// +kubebuilder:printcolumn:name="ResourceVersion",type="integer",JSONPath=".spec.metadata.resourceVersion"
// +kubebuilder:printcolumn:name="Created",type="date",JSONPath=".spec.metadata.createdAt"
// +kubebuilder:printcolumn:name="Modified",type="date",JSONPath=".spec.metadata.modifiedAt"

// VpDeploymentTarget is the Schema for the vpdeploymenttargets API
type VpDeploymentTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec VpDeploymentTargetObjectSpec `json:"spec,omitempty"`
}

// TODO: think about adding a field / status that keeps track of all the deployments with this target

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
