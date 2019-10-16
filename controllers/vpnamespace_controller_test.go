package controllers

import (
	"context"
	"time"

	"github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/appmanager-api-client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("VpNamespace Controller", func() {
	var reconciler VpNamespaceReconciler

	BeforeEach(func() {
		vpAPIClient := vpAPI.APIClient{}

		reconciler = VpNamespaceReconciler{
			Client:      k8sClient,
			Log:         logger,
			VPAPIClient: &vpAPIClient,
		}
	})

	Describe("updateResource", func() {
		var (
			key              types.NamespacedName
			created, fetched *v1beta1.VpNamespace
		)

		BeforeEach(func() {
			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}
			created = &v1beta1.VpNamespace{
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

		It("should update a k8s namespace with a VP namespace", func() {
			namespace := &vpAPI.Namespace{
				Kind:       "Namespace",
				ApiVersion: "v1",
				Metadata: &vpAPI.NamespaceMetadata{
					Id:              "2da2f867-5899-4bef-8ad0-9771bbac38b4",
					Name:            created.Name,
					CreatedAt:       time.Now(),
					ModifiedAt:      time.Now(),
					ResourceVersion: 1,
				},
				Status: &vpAPI.NamespaceStatus{State: "ACTIVE"},
			}

			Expect(reconciler.updateResource(created, namespace)).To(Succeed())

			fetched = &v1beta1.VpNamespace{}
			Expect(k8sClient.Get(context.TODO(), key, fetched)).To(Succeed())
			Expect(string(fetched.Status.State)).To(Equal(namespace.Status.State))
			Expect(fetched.Spec.Metadata.ResourceVersion).To(Equal(namespace.Metadata.ResourceVersion))
			Expect(fetched.Spec.Metadata.ID).To(Equal(namespace.Metadata.Id))
			Expect(fetched.ObjectMeta.Name).To(Equal(namespace.Metadata.Name))
		})
	})
})
