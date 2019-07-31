package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AddFinalizerToObjectMeta adds the finalizer to a K8s resource's object metadata
func AddFinalizerToObjectMeta(meta *metav1.ObjectMeta) bool {
	if ContainsString(meta.Finalizers, FinalizerName) {
		return false
	}
	meta.Finalizers = append(meta.Finalizers, FinalizerName)
	return true
}
