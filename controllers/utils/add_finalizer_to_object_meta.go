package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func AddFinalizerToObjectMeta(meta *metav1.ObjectMeta) bool {
	if ContainsString(meta.Finalizers, FinalizerName) {
		return false
	}
	meta.Finalizers = append(meta.Finalizers, FinalizerName)
	return true
}