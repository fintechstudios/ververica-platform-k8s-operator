package controllers

import (
	"context"
	"time"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("VpDeploymentTarget Controller", func() {
	var reconciler VpDeploymentTargetReconciler

	BeforeEach(func() {
		vpAPIClient := vpAPI.APIClient{}

		reconciler = VpDeploymentTargetReconciler{
			Client:      k8sClient,
			Log:         logger,
			VPAPIClient: &vpAPIClient,
		}
	})

	Describe("updateResource", func() {
		var (
			key              types.NamespacedName
			created, fetched *ververicaplatformv1beta1.VpDeploymentTarget
		)

		BeforeEach(func() {
			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}
			created = &ververicaplatformv1beta1.VpDeploymentTarget{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
			}
			Expect(k8sClient.Create(context.TODO(), created)).To(Succeed())
		})

		AfterEach(func() {
			Expect(k8sClient.Delete(context.TODO(), created)).To(Succeed())
		})

		It("should update a k8s deployment target with a VP deployment target", func() {
			depTarget := &vpAPI.DeploymentTarget{
				Kind:       "DeploymentTarget",
				ApiVersion: "v1",
				Metadata: &vpAPI.DeploymentTargetMetadata{
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
				Spec: &vpAPI.DeploymentTargetSpec{
					Kubernetes: &vpAPI.KubernetesTarget{
						Namespace: "default",
					},
					DeploymentPatchSet: []vpAPI.JsonPatchGeneric{
						{
							Op:    "add",
							Path:  "/test/field",
							Value: "data",
						},
						{
							Op:   "move",
							From: "/test/field",
							Path: "/test/field2",
						},
					},
				},
			}

			Expect(reconciler.updateResource(created, depTarget)).To(Succeed())

			fetched = &ververicaplatformv1beta1.VpDeploymentTarget{}
			Expect(k8sClient.Get(context.TODO(), key, fetched)).To(Succeed())
			Expect(fetched.Spec.Metadata.ResourceVersion).To(Equal(depTarget.Metadata.ResourceVersion))
			Expect(fetched.Spec.Metadata.ID).To(Equal(depTarget.Metadata.Id))
			Expect(fetched.Spec.Metadata.Labels).To(Equal(depTarget.Metadata.Labels))
			Expect(fetched.Spec.Metadata.Annotations).To(Equal(depTarget.Metadata.Annotations))
			Expect(fetched.Spec.Spec.DeploymentPatchSet).To(HaveLen(len(depTarget.Spec.DeploymentPatchSet)))
			for i, patch := range fetched.Spec.Spec.DeploymentPatchSet {
				depPatch := depTarget.Spec.DeploymentPatchSet[i]
				Expect(patch.From).To(Equal(depPatch.From))
				Expect(patch.Op).To(Equal(depPatch.Op))
				Expect(patch.Path).To(Equal(depPatch.Path))

				if depPatch.Value == nil {
					Expect(patch.Value).To(BeNil())
				} else {
					Expect(*patch.Value).To(Equal(depPatch.Value))
				}
			}
			Expect(fetched.Spec.Spec.Kubernetes.Namespace).To(Equal(depTarget.Spec.Kubernetes.Namespace))
			Expect(fetched.ObjectMeta.Name).To(Equal(depTarget.Metadata.Name))
		})
	})
})
