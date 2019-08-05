package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AddFinalizer adds the finalizer to a K8s resource's object metadata
func AddFinalizer(meta *metav1.ObjectMeta) bool {
	if ContainsString(meta.GetFinalizers(), FinalizerName) {
		return false
	}
	meta.Finalizers = append(meta.GetFinalizers(), FinalizerName)
	return true
}
