package nativeconverters

import (
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
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
							Spec: &v1beta2.VpDeploymentTemplateSpec{
								Kubernetes: &v1beta2.VpKubernetesOptions{
									Pods: &v1beta2.VpPodSpec{
										EnvVars: []core.EnvVar{
											{
												Name:  "TEST_ENV",
												Value: "TEST_VALUE",
											},
											{
												Name: "API_KEY",
												ValueFrom: &core.EnvVarSource{
													SecretKeyRef: &core.SecretKeySelector{
														LocalObjectReference: core.LocalObjectReference{
															Name: "some-secret",
														},
														Key:      "api-key",
														Optional: pointer.BoolPtr(true),
													},
												},
											},
										},
									},
								},
							},
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

			podSpec := dep.Spec.Template.Spec.Kubernetes.Pods
			Expect(podSpec.EnvVars).To(HaveLen(2))
			Expect(podSpec.EnvVars[1].ValueFrom.SecretKeyRef).ToNot(BeNil())
			Expect(podSpec.EnvVars[1].ValueFrom.SecretKeyRef.Name).To(Equal("some-secret"))
			Expect(podSpec.EnvVars[1].ValueFrom.SecretKeyRef.Key).To(Equal("api-key"))
			Expect(podSpec.EnvVars[1].ValueFrom.SecretKeyRef.Optional).ToNot(BeNil())
			Expect(*podSpec.EnvVars[1].ValueFrom.SecretKeyRef.Optional).To(BeTrue())

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
