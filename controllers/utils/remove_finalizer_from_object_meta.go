package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func RemoveFinalizerFromObjectMeta(meta *metav1.ObjectMeta) bool {
	if !ContainsString(meta.Finalizers, FinalizerName) {
		return false
	}
	meta.Finalizers = RemoveString(meta.Finalizers, FinalizerName)
	return true
}
