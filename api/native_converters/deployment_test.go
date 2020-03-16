package native_converters

import (
	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Deployment", func() {
	Describe("DeploymentFromNative", func() {
		var vpDeployment ververicaplatformv1beta1.VpDeployment

		BeforeEach(func() {
			vpDeployment = ververicaplatformv1beta1.VpDeployment{
				Spec: ververicaplatformv1beta1.VpDeploymentObjectSpec{
					Metadata: ververicaplatformv1beta1.VpMetadata{},
					Spec: ververicaplatformv1beta1.VpDeploymentSpec{
						Template: &ververicaplatformv1beta1.VpDeploymentTemplate{
							Metadata: &ververicaplatformv1beta1.VpDeploymentTemplateMetadata{},
							Spec:     &ververicaplatformv1beta1.VpDeploymentTemplateSpec{},
						},
					},
				},
				Status: ververicaplatformv1beta1.VpDeploymentStatus{
					State: ververicaplatformv1beta1.RunningState,
				},
			}
		})

		It("should map a K8s native deployment to the API respresentation", func() {
			dep, err := DeploymentFromNative(vpDeployment)
			Expect(err).ToNot(HaveOccurred())
			Expect(dep.Spec).ToNot(BeNil())
			Expect(dep.Metadata).ToNot(BeNil())
			Expect(dep.Status).ToNot(BeNil())
			Expect(dep.Status.State).To(Equal(string(ververicaplatformv1beta1.RunningState)))
		})

		Measure("conversion speed", func(b Benchmarker) {
			b.Time("from native", func() {
				_, _ = DeploymentFromNative(vpDeployment)
			})
		}, 10)
	})
})
