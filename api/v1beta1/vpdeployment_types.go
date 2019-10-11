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

// VpDeploymentStartFromSavepoint describes which savepoint, if any, to start the job with
type VpDeploymentStartFromSavepoint struct {
	// +optional
	Kind string `json:"kind,omitempty"`
}

// VpDeploymentTemplateMetadata
type VpDeploymentTemplateMetadata struct {
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
}

// VpArtifact describes a jar to run along with the Flink requirements
type VpArtifact struct {
	Kind string `json:"kind"`

	JarUri string `json:"jarUri"`
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
	Cpu resource.Quantity `json:"cpu,omitempty"`
	// +optional
	// +kubebuilder:validation:minLength=2
	Memory *string `json:"memory,omitempty"`
}

// VpLogging configures various loggers
type VpLogging struct {
	// +optional
	Log4jLoggers map[string]string `json:"log4jLoggers,omitempty"`
}

type VpJsonNode struct {
	// +optional
	ValueNode bool `json:"valueNode,omitempty"`
	// +optional
	ContainerNode bool `json:"containerNode,omitempty"`
	// +optional
	MissingNode bool `json:"missingNode,omitempty"`
	// +optional
	Object bool `json:"object,omitempty"`
	// +optional
	NodeType string `json:"nodeType,omitempty"`
	// +optional
	Pojo bool `json:"pojo,omitempty"`
	// +optional
	Number bool `json:"number,omitempty"`
	// +optional
	IntegralNumber bool `json:"integralNumber,omitempty"`
	// +optional
	FloatingPointNumber bool `json:"floatingPointNumber,omitempty"`
	// +optional
	Short bool `json:"short,omitempty"`
	// +optional
	Int_ bool `json:"int,omitempty"`
	// +optional
	Long bool `json:"long,omitempty"`
	// +optional
	Float bool `json:"float,omitempty"`
	// +optional
	Double bool `json:"double,omitempty"`
	// +optional
	BigDecimal bool `json:"bigDecimal,omitempty"`
	// +optional
	BigInteger bool `json:"bigInteger,omitempty"`
	// +optional
	Textual bool `json:"textual,omitempty"`
	// +optional
	Boolean bool `json:"boolean,omitempty"`
	// +optional
	Binary bool `json:"binary,omitempty"`
	// +optional
	Array bool `json:"array,omitempty"`
	// +optional
	Null bool `json:"null,omitempty"`
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

// DeploymentState is the enum of all possible deployment states
// Only one of the following states may be specified.
// +kubebuilder:validation:Enum=CANCELLED;RUNNING;TRANSITIONING;SUSPENDED;FAILED
type DeploymentState string

// All the allowed DeploymentStates
const (
	CancelledState     = DeploymentState("CANCELLED") // non-US spelling intentional
	RunningState       = DeploymentState("RUNNING")
	TransitioningState = DeploymentState("TRANSITIONING")
	SuspendedState     = DeploymentState("SUSPENDED")
	FailedState        = DeploymentState("FAILED")
)

// VpDeploymentSpec is the spec in the Ververica Platform
type VpDeploymentSpec struct {
	State DeploymentState `json:"state"`

	UpgradeStrategy *VpDeploymentUpgradeStrategy `json:"upgradeStrategy"`
	// +optional
	RestoreStrategy *VpDeploymentRestoreStrategy `json:"restoreStrategy,omitempty"`
	// +optional
	StartFromSavepoint *VpDeploymentStartFromSavepoint `json:"startFromSavepoint,omitempty"`
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
	State DeploymentState `json:"state,omitempty"`

	// +optional
	ResourceVersion int32 `json:"resourceVersion,omitempty"`

	// TODO: think about adding other information here, ie:
	// 		- list of Jobs
	//		- list of Events
	//		- list of K8s Pods created
	//		- potentially all dynamic data (id, etc.) ?
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Id",type="string",JSONPath=".spec.metadata.id"
// +kubebuilder:printcolumn:name="Created",type="date",JSONPath=".spec.metadata.createdAt"
// +kubebuilder:printcolumn:name="Modified",type="date",JSONPath=".spec.metadata.modifiedAt"
// +kubebuilder:printcolumn:name="Flink-Version",type="string",JSONPath=".spec.spec.template.spec.artifact.flinkVersion"
// +kubebuilder:printcolumn:name="Flink-Image-Tag",type="string",JSONPath=".spec.spec.template.spec.artifact.flinkImageTag"
// +kubebuilder:printcolumn:name="Flink-Image-Registry",type="string",JSONPath=".spec.spec.template.spec.artifact.flinkImageRegistry"
// +kubebuilder:printcolumn:name="Flink-Image-Repository",type="string",JSONPath=".spec.spec.template.spec.artifact.flinkImageRepository"

// VpDeployment is the Schema for the vpdeployments API
type VpDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   VpDeploymentObjectSpec `json:"spec"`
	Status VpDeploymentStatus     `json:"status,omitempty"`
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
