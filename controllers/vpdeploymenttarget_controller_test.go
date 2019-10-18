package controllers

import (
	"context"
	"fmt"
	"time"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/appmanager-api-client"
	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/annotations"
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
			Client:              k8sClient,
			Log:                 logger,
			AppManagerApiClient: &vpAPIClient,
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
					Annotations: make(map[string]string),
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
			Expect(annotations.Get(fetched.Annotations, annotations.ResourceVersion)).To(Equal(fmt.Sprint(depTarget.Metadata.ResourceVersion)))
			Expect(annotations.Get(fetched.Annotations, annotations.ID)).To(Equal(depTarget.Metadata.Id))
		})
	})
})
