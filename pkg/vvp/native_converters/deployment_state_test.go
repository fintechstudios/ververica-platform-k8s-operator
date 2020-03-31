package nativeconverters

import (
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("VpDeploymentState", func() {
	var deploymentStates = []v1beta2.VpDeploymentState{
		v1beta2.CancelledState,
		v1beta2.RunningState,
		v1beta2.TransitioningState,
		v1beta2.SuspendedState,
		v1beta2.FailedState,
		v1beta2.FinishedState,
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
			_, err := DeploymentStateFromNative(v1beta2.VpDeploymentState("not-a-state"))
			Expect(err).To(HaveOccurred())
		})
	})
})
