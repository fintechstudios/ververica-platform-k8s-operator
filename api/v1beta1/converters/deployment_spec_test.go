package converters

import (
	"reflect"

	"k8s.io/apimachinery/pkg/api/resource"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeploymentSpec", func() {
	var deploymentStateStr = "RUNNING"
	var deploymentState = ververicaplatformv1beta1.RunningState
	var deploymentTargetID = "2dd1ded8-eedb-4064-ba9a-006740d0f87a"
	var deploymentMaxSavepointAttempts = int32(4)
	var deploymentMaxCreationAttempts = int32(2)
	var deploymentUpgradeStrategy = "STATELESS"
	var deploymentRestoreStrategy = "NONE"
	var deploymentStartFromSavepoint = "s3://flink/a-savepoint"
	var deploymentRestoreAllowNonRestored = false
	var deploymentParallelism = int32(1)
	var deploymentNumTaskManagers = int32(2)
	var artifactKind = "JAR"
	var artifactJarURI = "s3://flink/a-jar"
	var artifactArgs = "--ed ed --and eddy"
	var artifactEntryClass = "com.fintechstudios.streaming"
	var artifactFlinkVersion = "1.8.0"
	var artifactFlinkRegistry = "registry.docker.com"
	var artifactFlinkRepository = "v1.4/flink"
	var artifactFlinkTag = "1.8.0_scala1.12"

	Describe("DeploymentSpecToNative", func() {
		var deploymentSpec vpAPI.DeploymentSpec
		var annotations map[string]string
		var flinkConfiguration map[string]string
		var resources map[string]vpAPI.ResourceSpec
		var vpResources map[string]ververicaplatformv1beta1.VpResourceSpec
		var log4jLoggers map[string]string
		var pods vpAPI.Pods

		BeforeEach(func() {
			annotations = map[string]string{
				"testing":           "true",
				"high-availability": "true",
			}
			resources = map[string]vpAPI.ResourceSpec{
				"jobmanager": {
					Cpu:    1.5,
					Memory: "2g",
				},
			}
			vpResources, _ = ResourcesToNative(resources)
			log4jLoggers = map[string]string{
				"":                   "DEBUG",
				"com.fintechstudios": "VERBOSE",
			}
			pods = vpAPI.Pods{}
			deploymentSpec = vpAPI.DeploymentSpec{
				State:                        deploymentStateStr,
				DeploymentTargetId:           deploymentTargetID,
				MaxJobCreationAttempts:       deploymentMaxCreationAttempts,
				MaxSavepointCreationAttempts: deploymentMaxSavepointAttempts,
				UpgradeStrategy: &vpAPI.DeploymentUpgradeStrategy{
					Kind: deploymentUpgradeStrategy,
				},
				RestoreStrategy: &vpAPI.DeploymentRestoreStrategy{
					Kind:                  deploymentRestoreStrategy,
					AllowNonRestoredState: deploymentRestoreAllowNonRestored,
				},
				Template: &vpAPI.DeploymentTemplate{
					Metadata: &vpAPI.DeploymentTemplateMetadata{
						Annotations: annotations,
					},
					Spec: &vpAPI.DeploymentTemplateSpec{
						Artifact: &vpAPI.Artifact{
							Kind:                 artifactKind,
							JarUri:               artifactJarURI,
							MainArgs:             artifactArgs,
							EntryClass:           artifactEntryClass,
							FlinkVersion:         artifactFlinkVersion,
							FlinkImageRegistry:   artifactFlinkRegistry,
							FlinkImageRepository: artifactFlinkRepository,
							FlinkImageTag:        artifactFlinkTag,
						},
						Parallelism:          deploymentParallelism,
						NumberOfTaskManagers: deploymentNumTaskManagers,
						Resources:            resources,
						FlinkConfiguration:   flinkConfiguration,
						Logging: &vpAPI.Logging{
							Log4jLoggers: log4jLoggers,
						},
						Kubernetes: &vpAPI.KubernetesOptions{
							Pods: &pods,
						},
					},
				},
			}
		})

		It("should map an API deployment spec to K8s native", func() {
			vpSpec, err := DeploymentSpecToNative(deploymentSpec)
			Expect(err).ToNot(HaveOccurred())
			Expect(vpSpec.StartFromSavepoint.Kind).To(Equal(deploymentStartFromSavepoint))

			Expect(vpSpec.RestoreStrategy.Kind).To(Equal(deploymentRestoreStrategy))
			Expect(vpSpec.RestoreStrategy.AllowNonRestoredState).To(Equal(deploymentRestoreAllowNonRestored))

			Expect(vpSpec.UpgradeStrategy.Kind).To(Equal(deploymentUpgradeStrategy))

			Expect(*vpSpec.MaxSavepointCreationAttempts).To(Equal(deploymentMaxSavepointAttempts))
			Expect(*vpSpec.MaxJobCreationAttempts).To(Equal(deploymentMaxCreationAttempts))
			Expect(vpSpec.DeploymentTargetID).To(Equal(deploymentTargetID))
			Expect(string(vpSpec.State)).To(Equal(deploymentStateStr))

			templateSpec := vpSpec.Template.Spec

			Expect(*templateSpec.Parallelism).To(Equal(deploymentParallelism))
			Expect(*templateSpec.NumberOfTaskManagers).To(Equal(deploymentNumTaskManagers))

			Expect(reflect.DeepEqual(templateSpec.Logging.Log4jLoggers, log4jLoggers)).To(BeTrue())
			Expect(reflect.DeepEqual(templateSpec.FlinkConfiguration, flinkConfiguration)).To(BeTrue())
			Expect(reflect.DeepEqual(templateSpec.Resources, vpResources)).To(BeTrue())

			artifact := templateSpec.Artifact
			Expect(artifact.FlinkImageTag).To(Equal(artifactFlinkTag))
			Expect(artifact.FlinkVersion).To(Equal(artifactFlinkVersion))
			Expect(artifact.FlinkImageRegistry).To(Equal(artifactFlinkRegistry))
			Expect(artifact.FlinkImageRepository).To(Equal(artifactFlinkRepository))
			Expect(artifact.EntryClass).To(Equal(artifactEntryClass))
			Expect(artifact.MainArgs).To(Equal(artifactArgs))
			Expect(artifact.Kind).To(Equal(artifactKind))
			Expect(artifact.JarURI).To(Equal(artifactJarURI))

			templateMetadata := vpSpec.Template.Metadata
			Expect(reflect.DeepEqual(templateMetadata.Annotations, annotations)).To(BeTrue())
		})

		It("should return an error given a deployment without a spec", func() {
			deploymentSpec.Template = nil
			_, err := DeploymentSpecToNative(deploymentSpec)
			Expect(err).To(HaveOccurred())
		})

		Measure("conversion speed", func(b Benchmarker) {
			b.Time("to native", func() {
				_, _ = DeploymentSpecToNative(deploymentSpec)
			})
		}, 10)
	})

	Describe("DeploymentSpecFromNative", func() {
		var vpDeploymentSpec ververicaplatformv1beta1.VpDeploymentSpec
		var annotations map[string]string
		var flinkConfiguration map[string]string
		var resources map[string]vpAPI.ResourceSpec
		var vpResources map[string]ververicaplatformv1beta1.VpResourceSpec
		var log4jLoggers map[string]string
		var vpPods ververicaplatformv1beta1.VpPodSpec

		BeforeEach(func() {
			annotations = map[string]string{
				"testing":           "true",
				"high-availability": "true",
			}
			mem := "2g"
			vpResources = map[string]ververicaplatformv1beta1.VpResourceSpec{
				"jobmanager": {
					CPU:    *resource.NewQuantity(2, resource.DecimalSI),
					Memory: &mem,
				},
			}
			resources, _ = ResourcesFromNative(vpResources)
			log4jLoggers = map[string]string{
				"":                   "DEBUG",
				"com.fintechstudios": "VERBOSE",
			}
			vpPods = ververicaplatformv1beta1.VpPodSpec{}
			vpDeploymentSpec = ververicaplatformv1beta1.VpDeploymentSpec{
				State:                        deploymentState,
				DeploymentTargetID:           deploymentTargetID,
				MaxJobCreationAttempts:       &deploymentMaxCreationAttempts,
				MaxSavepointCreationAttempts: &deploymentMaxSavepointAttempts,
				UpgradeStrategy: &ververicaplatformv1beta1.VpDeploymentUpgradeStrategy{
					Kind: deploymentUpgradeStrategy,
				},
				RestoreStrategy: &ververicaplatformv1beta1.VpDeploymentRestoreStrategy{
					Kind:                  deploymentRestoreStrategy,
					AllowNonRestoredState: deploymentRestoreAllowNonRestored,
				},
				StartFromSavepoint: &ververicaplatformv1beta1.VpDeploymentStartFromSavepoint{
					Kind: deploymentStartFromSavepoint,
				},
				Template: &ververicaplatformv1beta1.VpDeploymentTemplate{
					Metadata: &ververicaplatformv1beta1.VpDeploymentTemplateMetadata{
						Annotations: annotations,
					},
					Spec: &ververicaplatformv1beta1.VpDeploymentTemplateSpec{
						Artifact: &ververicaplatformv1beta1.VpArtifact{
							Kind:                 artifactKind,
							JarURI:               artifactJarURI,
							MainArgs:             artifactArgs,
							EntryClass:           artifactEntryClass,
							FlinkVersion:         artifactFlinkVersion,
							FlinkImageRegistry:   artifactFlinkRegistry,
							FlinkImageRepository: artifactFlinkRepository,
							FlinkImageTag:        artifactFlinkTag,
						},
						Parallelism:          &deploymentParallelism,
						NumberOfTaskManagers: &deploymentNumTaskManagers,
						Resources:            vpResources,
						FlinkConfiguration:   flinkConfiguration,
						Logging: &ververicaplatformv1beta1.VpLogging{
							Log4jLoggers: log4jLoggers,
						},
						Kubernetes: &ververicaplatformv1beta1.VpKubernetesOptions{
							Pods: &vpPods,
						},
					},
				},
			}
		})

		It("should map an API deployment spec to K8s native", func() {
			spec, err := DeploymentSpecFromNative(vpDeploymentSpec)
			Expect(err).ToNot(HaveOccurred())

			Expect(spec.RestoreStrategy.Kind).To(Equal(deploymentRestoreStrategy))
			Expect(spec.RestoreStrategy.AllowNonRestoredState).To(Equal(deploymentRestoreAllowNonRestored))

			Expect(spec.UpgradeStrategy.Kind).To(Equal(deploymentUpgradeStrategy))

			Expect(spec.DeploymentTargetId).To(Equal(deploymentTargetID))
			Expect(spec.MaxSavepointCreationAttempts).To(Equal(deploymentMaxSavepointAttempts))
			Expect(spec.MaxJobCreationAttempts).To(Equal(deploymentMaxCreationAttempts))
			Expect(spec.State).To(Equal(deploymentStateStr))

			templateSpec := spec.Template.Spec

			Expect(templateSpec.Parallelism).To(Equal(deploymentParallelism))
			Expect(templateSpec.NumberOfTaskManagers).To(Equal(deploymentNumTaskManagers))

			Expect(reflect.DeepEqual(templateSpec.Logging.Log4jLoggers, log4jLoggers)).To(BeTrue())
			Expect(reflect.DeepEqual(templateSpec.FlinkConfiguration, flinkConfiguration)).To(BeTrue())
			Expect(reflect.DeepEqual(templateSpec.Resources, resources)).To(BeTrue())

			artifact := templateSpec.Artifact
			Expect(artifact.FlinkImageTag).To(Equal(artifactFlinkTag))
			Expect(artifact.FlinkVersion).To(Equal(artifactFlinkVersion))
			Expect(artifact.FlinkImageRegistry).To(Equal(artifactFlinkRegistry))
			Expect(artifact.FlinkImageRepository).To(Equal(artifactFlinkRepository))
			Expect(artifact.EntryClass).To(Equal(artifactEntryClass))
			Expect(artifact.MainArgs).To(Equal(artifactArgs))
			Expect(artifact.Kind).To(Equal(artifactKind))
			Expect(artifact.JarUri).To(Equal(artifactJarURI))

			templateMetadata := spec.Template.Metadata
			Expect(reflect.DeepEqual(templateMetadata.Annotations, annotations)).To(BeTrue())
		})

		It("should return an error given a deployment without a spec", func() {
			vpDeploymentSpec.Template = nil
			_, err := DeploymentSpecFromNative(vpDeploymentSpec)
			Expect(err).To(HaveOccurred())
		})

		Measure("conversion speed", func(b Benchmarker) {
			b.Time("from native", func() {
				_, _ = DeploymentSpecFromNative(vpDeploymentSpec)
			})
		}, 10)
	})
})
