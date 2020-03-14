package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RemoveFinalizer removes the finalizer from a K8s object metadata
func RemoveFinalizer(meta *metav1.ObjectMeta) bool {
	if !ContainsString(meta.GetFinalizers(), FinalizerName) {
		return false
	}
	meta.SetFinalizers(RemoveString(meta.GetFinalizers(), FinalizerName))
	return true
}
