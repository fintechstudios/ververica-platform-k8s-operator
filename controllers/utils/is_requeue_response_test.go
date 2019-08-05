package utils

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ctrl "sigs.k8s.io/controller-runtime"
)

var _ = Describe("IsRequeueResponse", func() {
	When("error is nil", func() {
		It("should return false when Requeue is false", func() {
			Expect(IsRequeueResponse(ctrl.Result{Requeue: false}, nil)).To(BeFalse())
		})

		It("should return true when Requeue is true", func() {
			Expect(IsRequeueResponse(ctrl.Result{Requeue: true}, nil)).To(BeTrue())
		})

		It("should return false when RequeueAfter is 0", func() {
			Expect(IsRequeueResponse(ctrl.Result{RequeueAfter: 0}, nil)).To(BeFalse())
		})

		It("should return true when RequeueAfter is greater than 0", func() {
			Expect(IsRequeueResponse(ctrl.Result{RequeueAfter: 100}, nil)).To(BeTrue())
		})
	})

	It("should return true when error is not nil", func() {
		Expect(IsRequeueResponse(ctrl.Result{}, errors.New("an error"))).To(BeTrue())
	})
})
