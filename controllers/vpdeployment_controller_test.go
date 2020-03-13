package controllers

import (
	"context"
	"time"

	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1/converters"
	"github.com/fintechstudios/ververica-platform-k8s-operator/internal/annotations"
	appManager "github.com/fintechstudios/ververica-platform-k8s-operator/internal/appmanager-api-client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("VpDeployment Controller", func() {
	var reconciler VpDeploymentReconciler

	BeforeEach(func() {
		vpAPIClient := appManager.APIClient{}

		reconciler = VpDeploymentReconciler{
			Client:              k8sClient,
			Log:                 logger,
			AppManagerAPIClient: &vpAPIClient,
		}
	})

	Describe("updateResource", func() {
		var (
			key              types.NamespacedName
			created, fetched *v1beta1.VpDeployment
		)

		BeforeEach(func() {
			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}
			created = &v1beta1.VpDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: v1beta1.VpDeploymentObjectSpec{
					Metadata: v1beta1.VpMetadata{},
					Spec: v1beta1.VpDeploymentSpec{
						UpgradeStrategy: &v1beta1.VpDeploymentUpgradeStrategy{
							Kind: "STATELESS",
						},
						State: v1beta1.CancelledState,
						Template: &v1beta1.VpDeploymentTemplate{
							Spec: &v1beta1.VpDeploymentTemplateSpec{
								Artifact: &v1beta1.VpArtifact{
									Kind:   "JAR",
									JarURI: "https://jars.com/peanut-butter",
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
			dep := &appManager.Deployment{
				Kind:       "Deployment",
				ApiVersion: "v1",
				Metadata: &appManager.DeploymentMetadata{
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
				Spec: &appManager.DeploymentSpec{
					DeploymentTargetId: "4wt2a128-5899-4bef-8ad0-9771bbac38b4",
					UpgradeStrategy: &appManager.DeploymentUpgradeStrategy{
						Kind: "STATELESS",
					},
					State: "RUNNING",
					Template: &appManager.DeploymentTemplate{
						Spec: &appManager.DeploymentTemplateSpec{
							Artifact: &appManager.Artifact{
								Kind:   "JAR",
								JarUri: "https://jars.com/peanut-butter",
							},
						},
					},
				},
				Status: &appManager.DeploymentStatus{State: "RUNNING"},
			}

			Expect(reconciler.updateResource(created, dep)).To(Succeed())

			fetched = &v1beta1.VpDeployment{}
			Expect(k8sClient.Get(context.TODO(), key, fetched)).To(Succeed())
			Expect(annotations.Get(fetched.Annotations, annotations.ID)).To(Equal(dep.Metadata.Id))
			statusState, _ := converters.DeploymentStateToNative(dep.Status.State)
			Expect(fetched.Status.State).To(Equal(statusState))
		})
	})
})
