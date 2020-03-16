package nativeconverters

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	appmanagerapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"
)

// ResourcesToNative maps Ververica Platform API resources to native K8s  ones
func ResourcesToNative(resources map[string]appmanagerapi.ResourceSpec) (map[string]v1beta2.VpResourceSpec, error) {
	vpResources := make(map[string]v1beta2.VpResourceSpec)
	for k, v := range resources {
		res := v1beta2.VpResourceSpec{}
		if len(v.Memory) > 0 {
			res.Memory = &v.Memory
		}
		res.CPU = resource.MustParse(fmt.Sprintf("%f", v.Cpu))
		vpResources[k] = res
	}
	return vpResources, nil
}

// ResourcesFromNative maps native K8s resources to Ververica Platform API ones
func ResourcesFromNative(vpResources map[string]v1beta2.VpResourceSpec) (map[string]appmanagerapi.ResourceSpec, error) {
	resources := make(map[string]appmanagerapi.ResourceSpec)
	for k, v := range vpResources {
		res := appmanagerapi.ResourceSpec{}
		if v.Memory != nil {
			res.Memory = *v.Memory
		}
		res.Cpu = float64(v.CPU.MilliValue()) / 1000 // convert back to a plain float
		resources[k] = res
	}
	return resources, nil
}
