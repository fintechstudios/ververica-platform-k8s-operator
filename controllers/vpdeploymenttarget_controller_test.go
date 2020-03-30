package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	mocks "github.com/fintechstudios/ververica-platform-k8s-operator/mocks/vvp/appmanager"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/annotations"
	appmanagerapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("VpDeploymentTarget Controller", func() {
	var reconciler VpDeploymentTargetReconciler

	BeforeEach(func() {
		client := &mocks.Client{}

		reconciler = VpDeploymentTargetReconciler{
			Client:           k8sClient,
			Log:              logger,
			AppManagerClient: client,
		}
	})

	Describe("updateResource", func() {
		var (
			key              types.NamespacedName
			created, fetched *v1beta2.VpDeploymentTarget
		)

		BeforeEach(func() {
			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}
			created = &v1beta2.VpDeploymentTarget{
				ObjectMeta: metav1.ObjectMeta{
					Name:        "foo",
					Namespace:   "default",
					Annotations: make(map[string]string),
				},
			}
			Expect(k8sClient.Create(context.TODO(), created)).To(Succeed())
		})

		AfterEach(func() {
			Expect(k8sClient.Delete(context.TODO(), created)).To(Succeed())
		})

		It("should update a k8s deployment target with a VP deployment target", func() {
			depTarget := &appmanagerapi.DeploymentTarget{
				Kind:       "DeploymentTarget",
				ApiVersion: "v1",
				Metadata: &appmanagerapi.DeploymentTargetMetadata{
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
				Spec: &appmanagerapi.DeploymentTargetSpec{
					Kubernetes: &appmanagerapi.KubernetesTarget{
						Namespace: "default",
					},
				},
			}

			Expect(reconciler.updateResource(created, depTarget)).To(Succeed())

			fetched = &v1beta2.VpDeploymentTarget{}
			Expect(k8sClient.Get(context.TODO(), key, fetched)).To(Succeed())
			Expect(annotations.Get(fetched.Annotations, annotations.ResourceVersion)).To(Equal(fmt.Sprint(depTarget.Metadata.ResourceVersion)))
			Expect(annotations.Get(fetched.Annotations, annotations.ID)).To(Equal(depTarget.Metadata.Id))
		})
	})
})
