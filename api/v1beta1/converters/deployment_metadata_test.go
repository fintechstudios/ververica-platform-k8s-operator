package converters

import (
	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"reflect"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeploymentMetadata", func() {
	const deploymentId = "9cfce163-e969-4d75-8847-0c4309fcfe99"
	const deploymentName = "test-deployment"
	const deploymentNamespace = "testing"

	Describe("DeploymentMetadataToNative", func() {
		var annotations map[string]string
		var labels map[string]string
		var createdAt time.Time
		var modifiedAt time.Time
		var metadata vpAPI.DeploymentMetadata

		BeforeEach(func() {
			createdAt = time.Now()
			modifiedAt = time.Now()
			annotations = map[string]string{
				"testing":           "true",
				"high-availability": "false",
			}
			labels = map[string]string{
				"excellent": "adventure",
			}
			metadata = vpAPI.DeploymentMetadata{
				Id:          deploymentId,
				Annotations: annotations,
				Labels:      labels,
				Name:        deploymentName,
				Namespace:   deploymentNamespace,
				CreatedAt:   createdAt,
				ModifiedAt:  modifiedAt,
			}
		})

		It("should map an API deployment metadata to K8s native", func() {
			vpMetadata, err := DeploymentMetadataToNative(metadata)
			Expect(err).ToNot(HaveOccurred())
			Expect(vpMetadata.Name).To(Equal(deploymentName))
			Expect(vpMetadata.Namespace).To(Equal(deploymentNamespace))
			Expect(vpMetadata.ID).To(Equal(deploymentId))
			createdAtTime := metav1.NewTime(createdAt)
			modifiedAtTime := metav1.NewTime(modifiedAt)
			Expect(vpMetadata.CreatedAt.Equal(&createdAtTime)).To(BeTrue())
			Expect(vpMetadata.ModifiedAt.Equal(&modifiedAtTime)).To(BeTrue())
			Expect(reflect.DeepEqual(vpMetadata.Labels, labels)).To(BeTrue())
			Expect(reflect.DeepEqual(vpMetadata.Annotations, annotations)).To(BeTrue())
		})

		Measure("conversion speed", func(b Benchmarker) {
			b.Time("to native", func() {
				_, _ = DeploymentMetadataToNative(metadata)
			})
		}, 10)
	})

	Describe("DeploymentMetadataFromNative", func() {
		var annotations map[string]string
		var labels map[string]string
		var createdAt metav1.Time
		var modifiedAt metav1.Time
		var vpMetadata ververicaplatformv1beta1.VpDeploymentMetadata

		BeforeEach(func() {
			createdAt = metav1.NewTime(time.Now())
			modifiedAt = metav1.NewTime(time.Now())
			annotations = map[string]string{
				"testing":           "true",
				"high-availability": "false",
			}
			labels = map[string]string{
				"excellent": "adventure",
			}
			vpMetadata = ververicaplatformv1beta1.VpDeploymentMetadata{
				ID:          deploymentId,
				Annotations: annotations,
				Labels:      labels,
				Name:        deploymentName,
				Namespace:   deploymentNamespace,
				CreatedAt:   &createdAt,
				ModifiedAt:  &modifiedAt,
			}
		})

		It("should map an API deployment metadata to K8s native", func() {
			metadata, err := DeploymentMetadataFromNative(vpMetadata)
			Expect(err).ToNot(HaveOccurred())
			Expect(metadata.Name).To(Equal(deploymentName))
			Expect(metadata.Namespace).To(Equal(deploymentNamespace))
			Expect(metadata.Id).To(Equal(deploymentId))
			Expect(metadata.CreatedAt.Equal(createdAt.Time)).To(BeTrue())
			Expect(metadata.ModifiedAt.Equal(modifiedAt.Time)).To(BeTrue())
			Expect(reflect.DeepEqual(metadata.Labels, labels)).To(BeTrue())
			Expect(reflect.DeepEqual(metadata.Annotations, annotations)).To(BeTrue())
		})

		Measure("conversion speed", func(b Benchmarker) {
			b.Time("from native", func() {
				_, _ = DeploymentMetadataFromNative(vpMetadata)
			})
		}, 10)
	})
})
