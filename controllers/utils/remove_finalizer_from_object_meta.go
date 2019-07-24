package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RemoveFinalizerFromObjectMeta removes the finalizer from a K8s object metadata
func RemoveFinalizerFromObjectMeta(meta *metav1.ObjectMeta) bool {
	if !ContainsString(meta.Finalizers, FinalizerName) {
		return false
	}
	meta.Finalizers = RemoveString(meta.Finalizers, FinalizerName)
	return true
}
