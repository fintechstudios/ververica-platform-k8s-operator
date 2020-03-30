package v1beta2

import (
	"encoding/json"
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/annotations"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// Changes
// - dropped support for StartFromSavepoint
// - added FINISHED Deployment State
// - added Labels to pod specs

const depTargetAnnotationBase = "v1beta2.VpDeploymentTarget"

// use annotations to store version details
var (
	annDepTargetPatchSet = annotations.NewAnnotationName(depTargetAnnotationBase + ".patch-set")
)

// ConvertTo converts this v1beta2 version to a v1beta1 "Hub" version
func (src *VpDeploymentTarget) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.VpDeploymentTarget)

	// base conversion
	dst.ObjectMeta = src.ObjectMeta
	// Spec
	dst.Spec.Metadata = v1beta1.VpMetadata{
		Namespace:   src.Spec.Metadata.Namespace,
		Labels:      src.Spec.Metadata.Labels,
		Annotations: src.Spec.Metadata.Annotations,
	}

	if annotations.Has(src.Annotations, annDepTargetPatchSet) {
		dst.Annotations = annotations.EnsureExist(dst.Annotations)
		// store as an annotation
		data := annotations.Get(src.Annotations, annDepTargetPatchSet)
		if err := json.Unmarshal([]byte(data), &dst.Spec.Spec.DeploymentPatchSet); err != nil {
			return err
		}
	}

	dst.Spec.Spec.Kubernetes = v1beta1.VpKubernetesTarget{
		Namespace: src.Spec.Spec.Kubernetes.Namespace,
	}

	return nil
}

// ConvertTo converts from the "Hub" version to a v1beta2 version
func (dst *VpDeploymentTarget) ConvertFrom(srcRaw conversion.Hub) error { // nolint:golint
	src := srcRaw.(*v1beta1.VpDeploymentTarget)

	// base conversion
	dst.ObjectMeta = src.ObjectMeta
	// Spec
	dst.Spec.Metadata = VpMetadata{
		Namespace:   src.Spec.Metadata.Namespace,
		Labels:      src.Spec.Metadata.Labels,
		Annotations: src.Spec.Metadata.Annotations,
	}

	if src.Spec.Spec.DeploymentPatchSet != nil {
		// store as an annotation
		data, err := json.Marshal(src.Spec.Spec.DeploymentPatchSet)
		if err != nil {
			return err
		}
		dst.Annotations = annotations.Set(dst.Annotations,
			annotations.Pair(annDepTargetPatchSet, string(data)))
	}

	dst.Spec.Spec.Kubernetes = VpKubernetesTarget{
		Namespace: src.Spec.Spec.Kubernetes.Namespace,
	}

	return nil
}
