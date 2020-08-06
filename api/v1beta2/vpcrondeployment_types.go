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

package v1beta2

import (
	batchv2alpha1 "k8s.io/api/batch/v2alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Port in the batchv2alpha1 ConcurrencyPolicy
// +kubebuilder:validation:Enum=Allow;Forbid;Replace
type ConcurrencyPolicy string

const (
	// AllowConcurrent allows CronJobs to run concurrently.
	AllowConcurrent = ConcurrencyPolicy(batchv2alpha1.AllowConcurrent)

	// ForbidConcurrent forbids concurrent runs, skipping next run if previous
	// hasn't finished yet.
	ForbidConcurrent = ConcurrencyPolicy(batchv2alpha1.ForbidConcurrent)

	// ReplaceConcurrent cancels currently running job and replaces it with a new one.
	ReplaceConcurrent = ConcurrencyPolicy(batchv2alpha1.ReplaceConcurrent)
)

// DeploymentTemplateSpec describes the data a VpDeployment should have when created from a template
type DeploymentTemplateSpec struct {
	// Standard object's metadata of the jobs created from this template.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the deployment.
	// +optional
	Spec VpDeploymentObjectSpec `json:"spec,omitempty"`
}

// VpCronDeploymentSpec defines the desired state of VpCronDeployment
type VpCronDeploymentSpec struct {
	// +kubebuilder:validation:MinLength=0
	// The schedule in Cron format, see: https://en.wikipedia.org/wiki/Cron.
	// Also supports the nonstandard macros: https://en.wikipedia.org/wiki/Cron#Nonstandard_predefined_scheduling_definitions
	Schedule string `json:"schedule"`

	// The template for all created deployments
	VpDeploymentTemplate VpDeploymentObjectSpec `json:"vpDeploymentTemplate"`

	// Optional deadline in seconds for starting the job if it misses scheduled
	// time for any reason.  Missed deployment executions will be counted as failed ones.
	// +optional
	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds,omitempty"`

	// Specifies how to treat concurrent executions of a Deployment.
	// Uses the type as the native CronJob.
	// Valid values are:
	// - "Allow" (default): allows VpCronDeployments to run concurrently;
	// - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet;
	// - "Replace": cancels currently running job and replaces it with a new one
	// +optional
	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrencyPolicy"`
	// +optional
	Suspend *bool `json:"suspend,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +optional
	SuccessfulDeploymentsHistoryLimit *int32 `json:"successfulDeploymentsHistoryLimit,omitempty"`
	// +kubebuilder:validation:Minimum=0
	// +optional
	FailedDeploymentsHistoryLimit *int32 `json:"failedDeploymentsHistoryLimit,omitempty"`
}

// VpCronDeploymentStatus defines the observed state of VpCronDeployment
type VpCronDeploymentStatus struct {
	// +optional
	LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty"`

	// A list of pointers to currently running deployments.
	// +optional
	Active []corev1.ObjectReference `json:"active,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Schedule",type="string",JSONPath=".spec.schedule"
// +kubebuilder:printcolumn:name="Suspend",type="boolean",JSONPath=".spec.suspend"
// +kubebuilder:printcolumn:name="Last Schedule",type="date",JSONPath=".status.lastScheduleTime"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// VpCronDeployment is the Schema for the vpcrondeployments API
type VpCronDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpCronDeploymentSpec   `json:"spec,omitempty"`
	Status VpCronDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpCronDeploymentList contains a list of VpCronDeployment
type VpCronDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpCronDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VpCronDeployment{}, &VpCronDeploymentList{})
}
