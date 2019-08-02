package converters

import (
	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/resource"
)

var _ = Describe("ResourcesToNative", func() {
	memory := "2g"
	cpu := 2.0
	resources := map[string]vpAPI.ResourceSpec{
		"jobmanager": vpAPI.ResourceSpec{
			Cpu:    cpu,
			Memory: memory,
		},
		"taskmanager": vpAPI.ResourceSpec{
			Cpu:    cpu,
			Memory: memory,
		},
	}

	It("should map a API resource to K8s native", func() {
		vpResources, err := ResourcesToNative(resources)
		Expect(err).ToNot(HaveOccurred())
		Expect(len(vpResources)).To(Equal(2))
		Expect(vpResources).To(HaveKey("jobmanager"))
		Expect(vpResources).To(HaveKey("taskmanager"))
		for _, resource := range vpResources {
			fmtCpu := resource.Cpu.MilliValue() / 1000
			Expect(float64(fmtCpu)).To(Equal(cpu))
			Expect(*resource.Memory).To(Equal(memory))
		}
	})
})

var _ = Describe("ResourcesFromNative", func() {
	memory := "2g"
	cpu := resource.MustParse("2.0")
	vpResources := map[string]ververicaplatformv1beta1.VpResourceSpec{
		"jobmanager": ververicaplatformv1beta1.VpResourceSpec{
			Cpu:    cpu,
			Memory: &memory,
		},
		"taskmanager": ververicaplatformv1beta1.VpResourceSpec{
			Cpu:    cpu,
			Memory: &memory,
		},
	}

	It("should map K8s native resources to API", func() {
		resources, err := ResourcesFromNative(vpResources)
		Expect(err).ToNot(HaveOccurred())
		for _, resource := range resources {
			fmtCpu := cpu.MilliValue() / 1000
			Expect(resource.Cpu).To(Equal(float64(fmtCpu)))
			Expect(resource.Memory).To(Equal(memory))
		}
	})
})
