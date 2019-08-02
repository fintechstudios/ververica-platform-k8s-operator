package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AddFinalizer", func() {
	It("should add finalizer to empty metadata", func() {
		objMeta := metav1.ObjectMeta{}
		res := AddFinalizer(&objMeta)
		Expect(res).To(BeTrue())
		Expect(objMeta.Finalizers).To(HaveLen(1))
		Expect(objMeta.Finalizers).To(ContainElement(FinalizerName))
	})

	It("should add finalizer to non-empty metadata", func() {
		otherFinalizer := "some-other-finalizer"
		objMeta := metav1.ObjectMeta{
			Finalizers: []string{otherFinalizer},
		}
		res := AddFinalizer(&objMeta)
		Expect(res).To(BeTrue())
		Expect(objMeta.Finalizers).To(HaveLen(2))
		Expect(objMeta.Finalizers).To(ContainElement(FinalizerName))
		Expect(objMeta.Finalizers).To(ContainElement(otherFinalizer))
	})

	It("should not add the finalizer if it already exists", func() {
		objMeta := metav1.ObjectMeta{
			Finalizers: []string{FinalizerName},
		}
		res := AddFinalizer(&objMeta)
		Expect(res).To(BeFalse())
		Expect(objMeta.Finalizers).To(HaveLen(1))
		Expect(objMeta.Finalizers).To(ContainElement(FinalizerName))
	})
})
