package utils

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("RemoveFinalizerFromObjectMeta", func() {
	otherFinalizer := "some-other-finalizer"

	It("should not remove anything when the finalizer is not present", func() {
		objMeta := metav1.ObjectMeta{
			Finalizers: []string{otherFinalizer},
		}
		res := RemoveFinalizerFromObjectMeta(&objMeta)
		Expect(res).To(BeFalse())
		Expect(objMeta.Finalizers).To(HaveLen(1))
		Expect(objMeta.Finalizers).To(ContainElement(otherFinalizer))
	})

	It("should remove the finalizer when present", func() {
		objMeta := metav1.ObjectMeta{
			Finalizers: []string{otherFinalizer, FinalizerName},
		}
		res := RemoveFinalizerFromObjectMeta(&objMeta)
		Expect(res).To(BeTrue())
		Expect(objMeta.Finalizers).To(HaveLen(1))
		Expect(objMeta.Finalizers).ToNot(ContainElement(FinalizerName))
	})
})
