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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VpDeploymentMetadata represents all metadata from the VP API
type VpDeploymentMetadata struct {
	// Set through K8s obj meta name
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Id string `json:"id,omitempty"`
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// +optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	// +optional
	ModifiedAt *metav1.Time `json:"modifiedAt,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// +optional
	ResourceVersion int32 `json:"resourceVersion,omitempty"`
}

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
	FlinkImageTag string `json:"flinkImageTag,omitempty"`
}

// VpResourceSpec represents the resource requirements for components like the job and task managers
type VpResourceSpec struct {
	// +optional
	Cpu resource.Quantity `json:"cpu,omitempty"`
	// +optional
	Memory *string `json:"memory,omitempty"`
}

type VpLogging struct {
	// +optional
	Log4jLoggers map[string]string `json:"log4jLoggers,omitempty"`
}

// VpJsonNode is odd - seems incorrect but TBD
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

type VpVolumeAndMount struct {
	// +optional
	Name *string `json:"name,omitempty"`
	// +optional
	Volume *VpJsonNode `json:"volume,omitempty"`
	// +optional
	VolumeMount *VpJsonNode `json:"volumeMount,omitempty"`
}

type VpEnvVar struct {
	// +optional
	Name *string `json:"name,omitempty"`
	// +optional
	Value *string `json:"value,omitempty"`
	// +optional
	ValueFrom *VpJsonNode `json:"valueFrom,omitempty"`
}

type VpLocalObjectReference struct {
	Name string `json:"name"`
}

// VpPods
type VpPods struct {
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// +optional
	SecurityContext *VpJsonNode `json:"securityContext,omitempty"`
	// +optional
	Affinity *VpJsonNode `json:"affinity,omitempty"`
	// +optional
	Tolerations []VpJsonNode `json:"tolerations,omitempty"`
	// +optional
	VolumeMounts []VpVolumeAndMount `json:"volumeMounts,omitempty"`
	// +optional
	EnvVars []VpEnvVar `json:"envVars,omitempty"`
	// +optional
	ImagePullSecrets []VpLocalObjectReference `json:"imagePullSecrets,omitempty"`
}

// VpKubernetesOptions
type VpKubernetesOptions struct {
	Pods *VpPods `json:"pods,omitempty"`
}

// VpDeploymentTemplateSpec
type VpDeploymentTemplateSpec struct {
	Artifact *VpArtifact `json:"artifact"`
	// +optional
	Parallelism *int32 `json:"parallelism,omitempty"`
	// +optional
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

// VpDeploymentTemplate
type VpDeploymentTemplate struct {
	// +optional
	Metadata *VpDeploymentTemplateMetadata `json:"metadata,omitempty"`

	Spec *VpDeploymentTemplateSpec `json:"spec"`
}


// DeploymentState
// Only one of the following states may be specified.
// +kubebuilder:validation:Enum=CANCELLED;RUNNING;TRANSITIONING;SUSPENDED;FAILED
type DeploymentState string

const (
	CancelledState DeploymentState = "CANCELLED" // non-US spelling
	RunningState DeploymentState = "RUNNING"
	TransitioningState DeploymentState = "TRANSITIONING"
	SuspendedState DeploymentState = "SUSPENDED"
	FailedState DeploymentState = "FAILED"
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
	DeploymentTargetId string `json:"deploymentTargetId,omitempty"`
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
	Metadata VpDeploymentMetadata `json:"metadata"`
	Spec VpDeploymentSpec `json:"spec"`

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
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Id",type="string",JSONPath=".spec.metadata.id"
// +kubebuilder:printcolumn:name="ResourceVersion",type="integer",JSONPath=".spec.metadata.resourceVersion"
// +kubebuilder:printcolumn:name="Created",type="date",JSONPath=".spec.metadata.createdAt"
// +kubebuilder:printcolumn:name="Modified",type="date",JSONPath=".spec.metadata.modifiedAt"
// +kubebuilder:printcolumn:name="State",type="date",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="FlinkVersion",type="string",JSONPath=".spec.spec.template.spec.artifact.flinkVersion"
// +kubebuilder:printcolumn:name="FlinkImageTag",type="string",JSONPath=".spec.spec.template.spec.artifact.flinkImageTag"
// +kubebuilder:printcolumn:name="FlinkImageRegistry",type="string",JSONPath=".spec.spec.template.spec.artifact.flinkImageRegistry"

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
