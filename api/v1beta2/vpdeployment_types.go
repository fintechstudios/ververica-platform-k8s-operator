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
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VpDeploymentUpgradeStrategy describes how to upgrade a job
type VpDeploymentUpgradeStrategy struct {
	// +optional
	Kind string `json:"kind,omitempty"`
}

// VpDeploymentRestoreStrategy describes how to restore a job
type VpDeploymentRestoreStrategy struct {
	// +optional
	Kind string `json:"kind,omitempty"`
	// +optional
	AllowNonRestoredState bool `json:"allowNonRestoredState,omitempty"`
}

// VpDeploymentTemplateMetadata
type VpDeploymentTemplateMetadata struct {
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
}

// VpArtifact describes a jar to run along with the Flink requirements
type VpArtifact struct {
	Kind string `json:"kind"`

	JarURI string `json:"jarUri"`
	// +optional
	MainArgs string `json:"mainArgs,omitempty"`
	// +optional
	EntryClass string `json:"entryClass,omitempty"`
	// +optional
	FlinkVersion string `json:"flinkVersion,omitempty"`
	// +optional
	FlinkImageRegistry string `json:"flinkImageRegistry,omitempty"`
	// +optional
	FlinkImageRepository string `json:"flinkImageRepository,omitempty"`
	// +optional
	FlinkImageTag string `json:"flinkImageTag,omitempty"`
}

// VpResourceSpec represents the resource requirements for components like the job and task managers
type VpResourceSpec struct {
	// +optional
	CPU resource.Quantity `json:"cpu,omitempty"`
	// +optional
	// +kubebuilder:validation:minLength=2
	Memory *string `json:"memory,omitempty"`
}

// VpLogging configures various loggers
type VpLogging struct {
	// +optional
	Log4jLoggers map[string]string `json:"log4jLoggers,omitempty"`
}

// VpVolumeAndMount is a wrapper around both core.Volume and core.VolumeMount
type VpVolumeAndMount struct {
	Name        string            `json:"name"`
	Volume      *core.Volume      `json:"volume"`
	VolumeMount *core.VolumeMount `json:"volumeMount"`
}

// VpPodSpec is a subset of core.PodSpec, with annotations, env => envVars, and volume mounts
type VpPodSpec struct {
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// List of environment variables to set in the container.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	EnvVars []core.EnvVar `json:"envVars,omitempty" patchStrategy:"merge" patchMergeKey:"name"`

	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// +optional
	VolumeMounts []VpVolumeAndMount `json:"volumeMounts,omitempty"`

	// NodeSelector is a selector which must be true for the pod to fit on a node.
	// Selector which must match a node's labels for the pod to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty" protobuf:"bytes,7,rep,name=nodeSelector"`
	// SecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	SecurityContext *core.PodSecurityContext `json:"securityContext,omitempty" protobuf:"bytes,14,opt,name=securityContext"`
	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.
	// If specified, these secrets will be passed to individual puller implementations for them to use. For example,
	// in the case of docker, only DockerConfig type secrets are honored.
	// More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,15,rep,name=imagePullSecrets"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity,omitempty" protobuf:"bytes,18,opt,name=affinity"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations,omitempty" protobuf:"bytes,22,opt,name=tolerations"`
}

// VpKubernetesOptions allows users to configure K8s pods for Deployments
type VpKubernetesOptions struct {
	Pods *VpPodSpec `json:"pods,omitempty"`
}

// VpDeploymentTemplateSpec is the base spec for Deployment jobs
type VpDeploymentTemplateSpec struct {
	Artifact *VpArtifact `json:"artifact"`
	// +optional
	// +kubebuilder:validation:Minimum=1
	Parallelism *int32 `json:"parallelism,omitempty"`
	// +optional
	// +kubebuilder:validation:Minimum=1
	NumberOfTaskManagers *int32 `json:"numberOfTaskManagers,omitempty"`
	// +optional
	Resources map[string]VpResourceSpec `json:"resources,omitempty"`
	// +optional
	FlinkConfiguration map[string]string `json:"flinkConfiguration,omitempty"`
	// +optional
	Logging *VpLogging `json:"logging,omitempty"`
	// +optional
	Kubernetes *VpKubernetesOptions `json:"kubernetes,omitempty"`
}

// VpDeploymentTemplate is the template for Deployment jobs
type VpDeploymentTemplate struct {
	// +optional
	Metadata *VpDeploymentTemplateMetadata `json:"metadata,omitempty"`

	Spec *VpDeploymentTemplateSpec `json:"spec"`
}

// VpDeploymentState is the enum of all possible deployment states
// Only one of the following states may be specified.
// +kubebuilder:validation:Enum=CANCELLED;RUNNING;TRANSITIONING;SUSPENDED;FAILED;FINISHED
type VpDeploymentState string

// All the allowed DeploymentStates
const (
	CancelledState     = VpDeploymentState("CANCELLED") // non-US spelling intentional
	RunningState       = VpDeploymentState("RUNNING")
	TransitioningState = VpDeploymentState("TRANSITIONING")
	SuspendedState     = VpDeploymentState("SUSPENDED")
	FailedState        = VpDeploymentState("FAILED")
	FinishedState      = VpDeploymentState("FINISHED")
)

// VpDeploymentRunningCondition provide more details
// about the state of deployment
type VpDeploymentRunningCondition struct {
	Type               string      `json:"type"`
	Status             string      `json:"status"`
	Message            string      `json:"message"`
	Reason             string      `json:"reason"`
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	LastUpdateTime     metav1.Time `json:"lastUpdateTime,omitempty"`
}

// VpDeploymentRunningStatus gives extra information about the
// status of the underlying Flink job
// see: https://docs.ververica.com/user_guide/deployments/index.html#running
type VpDeploymentRunningStatus struct {
	// +optional
	Conditions []VpDeploymentRunningCondition `json:"conditions,omitempty"`
	JobID      string                         `json:"jobId"`
	// +optional
	TransitionTime metav1.Time `json:"transitionTime,omitempty"`
}

// VpDeploymentSpec is the spec in the Ververica Platform
type VpDeploymentSpec struct {
	State VpDeploymentState `json:"state"`

	UpgradeStrategy *VpDeploymentUpgradeStrategy `json:"upgradeStrategy"`
	// +optional
	RestoreStrategy *VpDeploymentRestoreStrategy `json:"restoreStrategy,omitempty"`
	// +optional
	DeploymentTargetID string `json:"deploymentTargetId,omitempty"`
	// +optional
	MaxSavepointCreationAttempts *int32 `json:"maxSavepointCreationAttempts,omitempty"`
	// +optional
	MaxJobCreationAttempts *int32 `json:"maxJobCreationAttempts,omitempty"`

	Template *VpDeploymentTemplate `json:"template"`
}

// VpDeploymentObjectSpec defines the desired state of VpDeployment
type VpDeploymentObjectSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// VP
	Metadata VpMetadata       `json:"metadata"`
	Spec     VpDeploymentSpec `json:"spec"`

	// DeploymentTargetName is an extension on the VP API
	// Must provide a spec.deploymentTargetId if not set
	// +optional
	DeploymentTargetName string `json:"deploymentTargetName,omitempty"`
}

// VpDeploymentStatus defines the observed state of VpDeployment
type VpDeploymentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	State VpDeploymentState `json:"state,omitempty"`

	// see: https://docs.ververica.com/user_guide/deployments/index.html#running
	// +optional
	Running *VpDeploymentRunningStatus `json:"running,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="Flink-Version",type="string",JSONPath=".spec.spec.template.spec.artifact.flinkVersion"
// +kubebuilder:printcolumn:name="Flink-Image-Tag",type="string",JSONPath=".spec.spec.template.spec.artifact.flinkImageTag"
// +kubebuilder:printcolumn:name="Flink-Image-Registry",type="string",JSONPath=".spec.spec.template.spec.artifact.flinkImageRegistry"
// +kubebuilder:printcolumn:name="Flink-Image-Repository",type="string",JSONPath=".spec.spec.template.spec.artifact.flinkImageRepository"

// VpDeployment is the Schema for the vpdeployments API
type VpDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec VpDeploymentObjectSpec `json:"spec"`
	// +optional
	Status *VpDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpDeploymentList contains a list of VpDeployment
type VpDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VpDeployment{}, &VpDeploymentList{})
}
