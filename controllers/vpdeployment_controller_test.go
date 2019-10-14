package controllers

import (
	"context"
	"time"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	"github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1/converters"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("VpDeployment Controller", func() {
	var reconciler VpDeploymentReconciler

	BeforeEach(func() {
		vpAPIClient := vpAPI.APIClient{}

		reconciler = VpDeploymentReconciler{
			Client:      k8sClient,
			Log:         logger,
			VPAPIClient: &vpAPIClient,
		}
	})

	Describe("updateResource", func() {
		var (
			key              types.NamespacedName
			created, fetched *ververicaplatformv1beta1.VpDeployment
		)

		BeforeEach(func() {
			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}
			created = &ververicaplatformv1beta1.VpDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: ververicaplatformv1beta1.VpDeploymentObjectSpec{
					Metadata: ververicaplatformv1beta1.VpMetadata{},
					Spec: ververicaplatformv1beta1.VpDeploymentSpec{
						UpgradeStrategy: &ververicaplatformv1beta1.VpDeploymentUpgradeStrategy{
							Kind: "STATELESS",
						},
						State: ververicaplatformv1beta1.CancelledState,
						Template: &ververicaplatformv1beta1.VpDeploymentTemplate{
							Spec: &ververicaplatformv1beta1.VpDeploymentTemplateSpec{
								Artifact: &ververicaplatformv1beta1.VpArtifact{
									Kind:   "JAR",
									JarUri: "https://jars.com/peanut-butter",
								},
							},
						},
					},
					DeploymentTargetName: "dep-target",
				},
			}
			Expect(k8sClient.Create(context.TODO(), created)).To(Succeed())
		})

		AfterEach(func() {
			Expect(k8sClient.Delete(context.TODO(), created)).To(Succeed())
		})

		It("should update a k8s deployment target with a VP deployment target", func() {
			dep := &vpAPI.Deployment{
				Kind:       "Deployment",
				ApiVersion: "v1",
				Metadata: &vpAPI.DeploymentMetadata{
					Id:              "2da2f867-5899-4bef-8ad0-9771bbac38b4",
					Name:            created.Name,
					CreatedAt:       time.Now(),
					ModifiedAt:      time.Now(),
					ResourceVersion: 1,
					Labels: map[string]string{
						"testing": "true",
					},
					Annotations: map[string]string{
						"non-production": "true",
					},
				},
				Spec: &vpAPI.DeploymentSpec{
					DeploymentTargetId: "4wt2a128-5899-4bef-8ad0-9771bbac38b4",
					UpgradeStrategy: &vpAPI.DeploymentUpgradeStrategy{
						Kind: "STATELESS",
					},
					State: "RUNNING",
					Template: &vpAPI.DeploymentTemplate{
						Spec: &vpAPI.DeploymentTemplateSpec{
							Artifact: &vpAPI.Artifact{
								Kind:   "JAR",
								JarUri: "https://jars.com/peanut-butter",
							},
						},
					},
				},
				Status: &vpAPI.DeploymentStatus{State: "RUNNING"},
			}

			Expect(reconciler.updateResource(created, dep)).To(Succeed())

			fetched = &ververicaplatformv1beta1.VpDeployment{}
			Expect(k8sClient.Get(context.TODO(), key, fetched)).To(Succeed())
			Expect(fetched.Spec.Metadata.ID).To(Equal(dep.Metadata.Id))
			Expect(fetched.Spec.Metadata.Labels).To(Equal(dep.Metadata.Labels))
			Expect(fetched.Spec.Metadata.Annotations).To(Equal(dep.Metadata.Annotations))
			Expect(fetched.ObjectMeta.Name).To(Equal(dep.Metadata.Name))
			Expect(fetched.Spec.DeploymentTargetName).To(Equal(created.Spec.DeploymentTargetName))
			state, _ := converters.DeploymentStateToNative(dep.Spec.State)
			Expect(fetched.Spec.Spec.State).To(Equal(state))
			statusState, _ := converters.DeploymentStateToNative(dep.Status.State)
			Expect(fetched.Status.State).To(Equal(statusState))
		})
	})
})
