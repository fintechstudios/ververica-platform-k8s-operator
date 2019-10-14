package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VpMetadata represents the base metadata for VP resources
type VpMetadata struct {
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	ID string `json:"id,omitempty"`
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
