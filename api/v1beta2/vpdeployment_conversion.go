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

const depAnnotationBase = "v1beta2.VpDeployment"

// use annotations to store version details
var (
	annDepStartFromSavepoint = annotations.NewAnnotationName(depAnnotationBase + ".start-from-savepoint")
	annDepState              = annotations.NewAnnotationName(depAnnotationBase + ".state")
	annDepStatusState        = annotations.NewAnnotationName(depAnnotationBase + ".status-state")
	annPodLabels             = annotations.NewAnnotationName(depAnnotationBase + ".pod-labels")
)

func convertToDeploymentState(state DeploymentState, annotation annotations.AnnotationName, notations map[string]string) v1beta1.DeploymentState {
	// If deployment status is FINISHED, mark it as RUNNING and add an annotation in v1beta1, otherwise, do a conversion
	if state == FinishedState {
		annotations.Set(
			notations,
			annotations.Pair(annotation, string(FinishedState)),
		)
		return v1beta1.RunningState
	}
	return v1beta1.DeploymentState(string(state))
}

func convertFromDeploymentState(state v1beta1.DeploymentState, annotation annotations.AnnotationName, notations map[string]string) DeploymentState {
	if annotations.Has(notations, annotation) {
		return DeploymentState(annotations.Get(notations, annotation))
	}

	return DeploymentState(string(state))
}

// ConvertTo converts this v1beta2 version to a v1beta1 "Hub" version
func (src *VpDeployment) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.VpDeployment)

	// base conversion
	dst.ObjectMeta = src.ObjectMeta
	// Spec
	dst.Spec.Metadata = v1beta1.VpMetadata{
		Namespace:   src.Spec.Metadata.Namespace,
		Labels:      src.Spec.Metadata.Labels,
		Annotations: src.Spec.Metadata.Annotations,
	}
	dst.Spec.DeploymentTargetName = src.Spec.DeploymentTargetName
	// Spec.Spec
	dst.Spec.Spec.DeploymentTargetID = src.Spec.Spec.DeploymentTargetID
	dst.Spec.Spec.MaxJobCreationAttempts = src.Spec.Spec.MaxJobCreationAttempts
	dst.Spec.Spec.MaxSavepointCreationAttempts = src.Spec.Spec.MaxSavepointCreationAttempts
	dst.Spec.Spec.UpgradeStrategy = &v1beta1.VpDeploymentUpgradeStrategy{
		Kind: src.Spec.Spec.UpgradeStrategy.Kind,
	}
	dst.Spec.Spec.State = convertToDeploymentState(src.Spec.Spec.State, annDepState, dst.Annotations)

	if src.Spec.Spec.RestoreStrategy != nil {
		dst.Spec.Spec.RestoreStrategy = &v1beta1.VpDeploymentRestoreStrategy{
			Kind:                  src.Spec.Spec.RestoreStrategy.Kind,
			AllowNonRestoredState: src.Spec.Spec.RestoreStrategy.AllowNonRestoredState,
		}
	}

	// StartFromSavepoint in v1beta2 is gone, unless stored in annotation
	if annotations.Has(src.Annotations, annDepStartFromSavepoint) {
		an := annotations.Get(src.Annotations, annDepStartFromSavepoint)
		if err := json.Unmarshal([]byte(an), dst.Spec.Spec.StartFromSavepoint); err != nil {
			return err
		}
	}

	// Spec.Spec template
	srcTmpl := src.Spec.Spec.Template
	dstTmpl := &v1beta1.VpDeploymentTemplate{}
	if srcTmpl.Metadata != nil {
		dstTmpl.Metadata = &v1beta1.VpDeploymentTemplateMetadata{
			Annotations: srcTmpl.Metadata.Annotations,
		}
	}

	dstTmpl.Spec.Artifact = &v1beta1.VpArtifact{
		Kind:                 srcTmpl.Spec.Artifact.Kind,
		JarURI:               srcTmpl.Spec.Artifact.JarURI,
		MainArgs:             srcTmpl.Spec.Artifact.MainArgs,
		EntryClass:           srcTmpl.Spec.Artifact.EntryClass,
		FlinkVersion:         srcTmpl.Spec.Artifact.FlinkVersion,
		FlinkImageRegistry:   srcTmpl.Spec.Artifact.FlinkImageRegistry,
		FlinkImageRepository: srcTmpl.Spec.Artifact.FlinkImageRepository,
		FlinkImageTag:        srcTmpl.Spec.Artifact.FlinkImageTag,
	}

	dstTmpl.Spec.FlinkConfiguration = srcTmpl.Spec.FlinkConfiguration

	if srcTmpl.Spec.Resources != nil {
		dstRes := dstTmpl.Spec.Resources
		dstRes = make(map[string]v1beta1.VpResourceSpec)
		for name, res := range srcTmpl.Spec.Resources {
			dstRes[name] = v1beta1.VpResourceSpec{
				CPU:    res.CPU,
				Memory: res.Memory,
			}
		}
	}

	if srcTmpl.Spec.Logging != nil {
		dstTmpl.Spec.Logging = &v1beta1.VpLogging{Log4jLoggers: srcTmpl.Spec.Logging.Log4jLoggers}
	}

	if srcTmpl.Spec.NumberOfTaskManagers != nil {
		dstTmpl.Spec.NumberOfTaskManagers = srcTmpl.Spec.NumberOfTaskManagers
	}

	if srcTmpl.Spec.Parallelism != nil {
		dstTmpl.Spec.Parallelism = srcTmpl.Spec.Parallelism
	}

	if srcTmpl.Spec.Kubernetes != nil && srcTmpl.Spec.Kubernetes.Pods != nil {
		srcPods := srcTmpl.Spec.Kubernetes.Pods

		// save the labels as an annotation
		labels, err := json.Marshal(srcPods.Labels)
		if err != nil {
			return err
		}
		annotations.Set(
			dst.Annotations,
			annotations.Pair(annPodLabels, string(labels)),
		)

		dstPods := &v1beta1.VpPodSpec{
			Annotations:      srcPods.Annotations,
			EnvVars:          srcPods.EnvVars,
			NodeSelector:     srcPods.NodeSelector,
			SecurityContext:  srcPods.SecurityContext,
			ImagePullSecrets: srcPods.ImagePullSecrets,
			Affinity:         srcPods.Affinity,
			Tolerations:      srcPods.Tolerations,
		}

		if srcPods.VolumeMounts != nil {
			dstPods.VolumeMounts = make([]v1beta1.VpVolumeAndMount, len(srcPods.VolumeMounts))
			for i, e := range srcPods.VolumeMounts {
				dstPods.VolumeMounts[i] = v1beta1.VpVolumeAndMount{
					Name:        e.Name,
					Volume:      e.Volume,
					VolumeMount: e.VolumeMount,
				}
			}
		}

		dstTmpl.Spec.Kubernetes = &v1beta1.VpKubernetesOptions{Pods: dstPods}
	}

	dst.Spec.Spec.Template = dstTmpl

	// Status
	dst.Status.State = convertToDeploymentState(src.Status.State, annDepStatusState, dst.Annotations)

	return nil
}

// ConvertTo converts from the "Hub" version to a v1beta2 version
func (dst *VpDeployment) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.VpDeployment)

	// base conversion
	dst.ObjectMeta = src.ObjectMeta
	// Spec
	dst.Spec.Metadata = VpMetadata{
		Namespace:   src.Spec.Metadata.Namespace,
		Labels:      src.Spec.Metadata.Labels,
		Annotations: src.Spec.Metadata.Annotations,
	}
	dst.Spec.DeploymentTargetName = src.Spec.DeploymentTargetName
	// Spec.Spec
	dst.Spec.Spec.DeploymentTargetID = src.Spec.Spec.DeploymentTargetID
	dst.Spec.Spec.MaxJobCreationAttempts = src.Spec.Spec.MaxJobCreationAttempts
	dst.Spec.Spec.MaxSavepointCreationAttempts = src.Spec.Spec.MaxSavepointCreationAttempts
	dst.Spec.Spec.UpgradeStrategy = &VpDeploymentUpgradeStrategy{
		Kind: src.Spec.Spec.UpgradeStrategy.Kind,
	}
	if src.Spec.Spec.RestoreStrategy != nil {
		dst.Spec.Spec.RestoreStrategy = &VpDeploymentRestoreStrategy{
			Kind:                  src.Spec.Spec.RestoreStrategy.Kind,
			AllowNonRestoredState: src.Spec.Spec.RestoreStrategy.AllowNonRestoredState,
		}
	}
	dst.Spec.Spec.State = convertFromDeploymentState(src.Spec.Spec.State, annDepState, dst.Annotations)

	if src.Spec.Spec.StartFromSavepoint != nil {
		data, err := json.Marshal(src.Spec.Spec.StartFromSavepoint)
		if err != nil {
			return err
		}
		annotations.Set(dst.Annotations,
			annotations.Pair(annDepStartFromSavepoint, string(data)))
	}

	// Spec.Spec template
	srcTmpl := src.Spec.Spec.Template
	dstTmpl := &VpDeploymentTemplate{}
	dstTmpl.Spec.Artifact = &VpArtifact{
		Kind:                 srcTmpl.Spec.Artifact.Kind,
		JarURI:               srcTmpl.Spec.Artifact.JarURI,
		MainArgs:             srcTmpl.Spec.Artifact.MainArgs,
		EntryClass:           srcTmpl.Spec.Artifact.EntryClass,
		FlinkVersion:         srcTmpl.Spec.Artifact.FlinkVersion,
		FlinkImageRegistry:   srcTmpl.Spec.Artifact.FlinkImageRegistry,
		FlinkImageRepository: srcTmpl.Spec.Artifact.FlinkImageRepository,
		FlinkImageTag:        srcTmpl.Spec.Artifact.FlinkImageTag,
	}

	dstTmpl.Spec.FlinkConfiguration = srcTmpl.Spec.FlinkConfiguration

	if srcTmpl.Spec.Resources != nil {
		dstRes := dstTmpl.Spec.Resources
		dstRes = make(map[string]VpResourceSpec)
		for name, res := range srcTmpl.Spec.Resources {
			dstRes[name] = VpResourceSpec{
				CPU:    res.CPU,
				Memory: res.Memory,
			}
		}
	}

	if srcTmpl.Spec.Logging != nil {
		dstTmpl.Spec.Logging = &VpLogging{Log4jLoggers: srcTmpl.Spec.Logging.Log4jLoggers}
	}

	if srcTmpl.Spec.NumberOfTaskManagers != nil {
		dstTmpl.Spec.NumberOfTaskManagers = srcTmpl.Spec.NumberOfTaskManagers
	}

	if srcTmpl.Spec.Parallelism != nil {
		dstTmpl.Spec.Parallelism = srcTmpl.Spec.Parallelism
	}

	if srcTmpl.Spec.Kubernetes != nil && srcTmpl.Spec.Kubernetes.Pods != nil {
		srcPods := srcTmpl.Spec.Kubernetes.Pods
		dstPods := &VpPodSpec{
			Annotations:      srcPods.Annotations,
			EnvVars:          srcPods.EnvVars,
			NodeSelector:     srcPods.NodeSelector,
			SecurityContext:  srcPods.SecurityContext,
			ImagePullSecrets: srcPods.ImagePullSecrets,
			Affinity:         srcPods.Affinity,
			Tolerations:      srcPods.Tolerations,
		}
		// unbundle stored labels
		if annotations.Has(src.Annotations, annPodLabels) {
			data := annotations.Get(src.Annotations, annPodLabels)
			if err := json.Unmarshal([]byte(data), &dstPods.Labels); err != nil {
				return err
			}
		}

		if srcPods.VolumeMounts != nil {
			dstPods.VolumeMounts = make([]VpVolumeAndMount, len(srcPods.VolumeMounts))
			for i, e := range srcPods.VolumeMounts {
				dstPods.VolumeMounts[i] = VpVolumeAndMount{
					Name:        e.Name,
					Volume:      e.Volume,
					VolumeMount: e.VolumeMount,
				}
			}
		}

		dstTmpl.Spec.Kubernetes = &VpKubernetesOptions{Pods: dstPods}
	}

	dst.Spec.Spec.Template = dstTmpl

	// Status
	dst.Status.State = convertFromDeploymentState(src.Status.State, annDepStatusState, dst.Annotations)

	return nil
}
