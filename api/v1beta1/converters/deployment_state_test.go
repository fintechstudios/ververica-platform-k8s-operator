package converters

import (
	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeploymentState", func() {
	var deploymentStates = []ververicaplatformv1beta1.DeploymentState{
		ververicaplatformv1beta1.CancelledState,
		ververicaplatformv1beta1.RunningState,
		ververicaplatformv1beta1.TransitioningState,
		ververicaplatformv1beta1.SuspendedState,
		ververicaplatformv1beta1.FailedState,
	}

	Describe("DeploymentStateToNative", func() {
		It("should map an API status to K8s native", func() {
			for _, state := range deploymentStates {
				mappedState, err := DeploymentStateToNative(string(state))
				Expect(err).ToNot(HaveOccurred())
				Expect(mappedState).To(Equal(state))
			}
		})

		It("should return an error given an invalid state", func() {
			_, err := DeploymentStateToNative("not-a-state")
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("DeploymentStateFromNative", func() {
		It("should map a K8s native status to  API", func() {
			for _, state := range deploymentStates {
				mappedState, err := DeploymentStateFromNative(state)
				Expect(err).ToNot(HaveOccurred())
				Expect(mappedState).To(Equal(string(state)))
			}
		})

		It("should return an error given an invalid state", func() {
			_, err := DeploymentStateFromNative(ververicaplatformv1beta1.DeploymentState("not-a-state"))
			Expect(err).To(HaveOccurred())
		})
	})
})
