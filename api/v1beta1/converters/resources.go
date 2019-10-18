package converters

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/appmanager-api-client"
)

// ResourcesToNative maps Ververica Platform API resources to native K8s  ones
func ResourcesToNative(resources map[string]vpAPI.ResourceSpec) (map[string]ververicaplatformv1beta1.VpResourceSpec, error) {
	vpResources := make(map[string]ververicaplatformv1beta1.VpResourceSpec)
	for k, v := range resources {
		res := ververicaplatformv1beta1.VpResourceSpec{}
		if len(v.Memory) > 0 {
			res.Memory = &v.Memory
		}
		res.Cpu = resource.MustParse(fmt.Sprintf("%f", v.Cpu))
		vpResources[k] = res
	}
	return vpResources, nil
}

// ResourcesFromNative maps native K8s resources to Ververica Platform API ones
func ResourcesFromNative(vpResources map[string]ververicaplatformv1beta1.VpResourceSpec) (map[string]vpAPI.ResourceSpec, error) {
	resources := make(map[string]vpAPI.ResourceSpec)
	for k, v := range vpResources {
		res := vpAPI.ResourceSpec{}
		if v.Memory != nil {
			res.Memory = *v.Memory
		}
		res.Cpu = float64(v.Cpu.MilliValue()) / 1000 // convert back to a plain float
		resources[k] = res
	}
	return resources, nil
}
