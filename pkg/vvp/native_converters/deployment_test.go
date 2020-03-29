package nativeconverters

import (
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Deployment", func() {
	Describe("DeploymentFromNative", func() {
		var vpDeployment v1beta2.VpDeployment

		BeforeEach(func() {
			vpDeployment = v1beta2.VpDeployment{
				Spec: v1beta2.VpDeploymentObjectSpec{
					Metadata: v1beta2.VpMetadata{},
					Spec: v1beta2.VpDeploymentSpec{
						Template: &v1beta2.VpDeploymentTemplate{
							Metadata: &v1beta2.VpDeploymentTemplateMetadata{},
							Spec:     &v1beta2.VpDeploymentTemplateSpec{},
						},
					},
				},
				Status: v1beta2.VpDeploymentStatus{
					State: v1beta2.RunningState,
				},
			}
		})

		It("should map a K8s native deployment to the API respresentation", func() {
			dep, err := DeploymentFromNative(vpDeployment)
			Expect(err).ToNot(HaveOccurred())
			Expect(dep.Spec).ToNot(BeNil())
			Expect(dep.Metadata).ToNot(BeNil())
			Expect(dep.Status).ToNot(BeNil())
			Expect(dep.Status.State).To(Equal(string(v1beta2.RunningState)))
		})

		Measure("conversion speed", func(b Benchmarker) {
			b.Time("from native", func() {
				_, _ = DeploymentFromNative(vpDeployment)
			})
		}, 10)
	})
})
