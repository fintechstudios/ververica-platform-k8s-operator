package nativeconverters

import (
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
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
				Status: &v1beta2.VpDeploymentStatus{
					State: v1beta2.RunningState,
					Running: &v1beta2.VpDeploymentRunningStatus{
						Conditions: []v1beta2.VpDeploymentRunningCondition{
							{
								Type:    "ClusterUnreachable",
								Message: "Unknown cluster failure",
								Status:  "Unknown",
								Reason:  "Failed to contact cluster",
								LastTransitionTime: metav1.NewTime(
									utils.MustParseTime(time.RFC3339, "2020-01-02T15:04:05Z"),
								),
								LastUpdateTime: metav1.NewTime(
									utils.MustParseTime(time.RFC3339, "2020-01-02T14:00:00Z"),
								),
							},
							{
								Type:    "JobFailing",
								Message: "Something is not happy",
								Status:  "True",
								Reason:  "error message here?",
								LastTransitionTime: metav1.NewTime(
									utils.MustParseTime(time.RFC3339, "2021-01-02T15:04:05Z"),
								),
								LastUpdateTime: metav1.NewTime(
									utils.MustParseTime(time.RFC3339, "2021-01-02T14:00:00Z"),
								),
							},
						},
						JobID: "a-job-id",
						TransitionTime: metav1.NewTime(
							utils.MustParseTime(time.RFC3339, "2020-01-02T15:04:05Z"),
						),
					},
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
			Expect(dep.Status.Running).ToNot(BeNil())
			Expect(dep.Status.Running.JobId).To(Equal("a-job-id"))
			Expect(dep.Status.Running.Conditions).To(HaveLen(2))
		})

		Measure("conversion speed", func(b Benchmarker) {
			b.Time("from native", func() {
				_, _ = DeploymentFromNative(vpDeployment)
			})
		}, 10)
	})
})
