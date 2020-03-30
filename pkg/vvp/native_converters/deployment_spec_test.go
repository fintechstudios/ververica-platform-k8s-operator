package nativeconverters

import (
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/resource"
	"reflect"
)

var _ = Describe("DeploymentSpec", func() {
	var deploymentStateStr = "RUNNING"
	var deploymentState = v1beta2.RunningState
	var deploymentTargetID = "2dd1ded8-eedb-4064-ba9a-006740d0f87a"
	var deploymentMaxSavepointAttempts = int32(4)
	var deploymentMaxCreationAttempts = int32(2)
	var deploymentUpgradeStrategy = "STATELESS"
	var deploymentRestoreStrategy = "NONE"
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
		var deploymentSpec appmanagerapi.DeploymentSpec
		var annotations map[string]string
		var flinkConfiguration map[string]string
		var resources map[string]appmanagerapi.ResourceSpec
		var vpResources map[string]v1beta2.VpResourceSpec
		var log4jLoggers map[string]string
		var pods appmanagerapi.Pods

		BeforeEach(func() {
			annotations = map[string]string{
				"testing":           "true",
				"high-availability": "true",
			}
			resources = map[string]appmanagerapi.ResourceSpec{
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
			pods = appmanagerapi.Pods{}
			deploymentSpec = appmanagerapi.DeploymentSpec{
				State:                        deploymentStateStr,
				DeploymentTargetId:           deploymentTargetID,
				MaxJobCreationAttempts:       deploymentMaxCreationAttempts,
				MaxSavepointCreationAttempts: deploymentMaxSavepointAttempts,
				UpgradeStrategy: &appmanagerapi.DeploymentUpgradeStrategy{
					Kind: deploymentUpgradeStrategy,
				},
				RestoreStrategy: &appmanagerapi.DeploymentRestoreStrategy{
					Kind:                  deploymentRestoreStrategy,
					AllowNonRestoredState: deploymentRestoreAllowNonRestored,
				},
				Template: &appmanagerapi.DeploymentTemplate{
					Metadata: &appmanagerapi.DeploymentTemplateMetadata{
						Annotations: annotations,
					},
					Spec: &appmanagerapi.DeploymentTemplateSpec{
						Artifact: &appmanagerapi.Artifact{
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
						Logging: &appmanagerapi.Logging{
							Log4jLoggers: log4jLoggers,
						},
						Kubernetes: &appmanagerapi.KubernetesOptions{
							Pods: &pods,
						},
					},
				},
			}
		})

		It("should map an API deployment spec to K8s native", func() {
			vpSpec, err := DeploymentSpecToNative(deploymentSpec)
			Expect(err).ToNot(HaveOccurred())

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
		var vpDeploymentSpec v1beta2.VpDeploymentSpec
		var annotations map[string]string
		var flinkConfiguration map[string]string
		var resources map[string]appmanagerapi.ResourceSpec
		var vpResources map[string]v1beta2.VpResourceSpec
		var log4jLoggers map[string]string
		var vpPods v1beta2.VpPodSpec

		BeforeEach(func() {
			annotations = map[string]string{
				"testing":           "true",
				"high-availability": "true",
			}
			mem := "2g"
			vpResources = map[string]v1beta2.VpResourceSpec{
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
			vpPods = v1beta2.VpPodSpec{}
			vpDeploymentSpec = v1beta2.VpDeploymentSpec{
				State:                        deploymentState,
				DeploymentTargetID:           deploymentTargetID,
				MaxJobCreationAttempts:       &deploymentMaxCreationAttempts,
				MaxSavepointCreationAttempts: &deploymentMaxSavepointAttempts,
				UpgradeStrategy: &v1beta2.VpDeploymentUpgradeStrategy{
					Kind: deploymentUpgradeStrategy,
				},
				RestoreStrategy: &v1beta2.VpDeploymentRestoreStrategy{
					Kind:                  deploymentRestoreStrategy,
					AllowNonRestoredState: deploymentRestoreAllowNonRestored,
				},
				Template: &v1beta2.VpDeploymentTemplate{
					Metadata: &v1beta2.VpDeploymentTemplateMetadata{
						Annotations: annotations,
					},
					Spec: &v1beta2.VpDeploymentTemplateSpec{
						Artifact: &v1beta2.VpArtifact{
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
						Logging: &v1beta2.VpLogging{
							Log4jLoggers: log4jLoggers,
						},
						Kubernetes: &v1beta2.VpKubernetesOptions{
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
