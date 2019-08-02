package converters

import (
	"reflect"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeploymentSpec", func() {
	const deploymentState = "RUNNING"
	const deploymentTargetID = "2dd1ded8-eedb-4064-ba9a-006740d0f87a"
	const deploymentMaxSavepointAttempts = int32(4)
	const deploymentMaxCreationAttempts = int32(2)
	const deploymentUpgradeStrategy = "STATELESS"
	const deploymentRestoreStrategy = "NONE"
	const deploymentStartFromSavepoint = "s3://flink/a-savepoint"
	const deploymentRestoreAllowNonRestored = false
	const deploymentParallelism = int32(1)
	const deploymentNumTaskManagers = int32(2)
	const artifactKind = "JAR"
	const artifactJarUri = "s3://flink/a-jar"
	const artifactArgs = "--ed ed --and eddy"
	const artifactEntryClass = "com.fintechstudios.streaming"
	const artifactFlinkVersion = "1.8.0"
	const artifactFlinkRegistry = "registry.docker.com"
	const artifactFlinkTag = "1.8.0_scala1.12"

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
				"jobmanager": vpAPI.ResourceSpec{
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
				State:                        deploymentState,
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
				StartFromSavepoint: &vpAPI.DeploymentStartFromSavepoint{
					Kind: deploymentStartFromSavepoint,
				},
				Template: &vpAPI.DeploymentTemplate{
					Metadata: &vpAPI.DeploymentTemplateMetadata{
						Annotations: annotations,
					},
					Spec: &vpAPI.DeploymentTemplateSpec{
						Artifact: &vpAPI.Artifact{
							Kind:               artifactKind,
							JarUri:             artifactJarUri,
							MainArgs:           artifactArgs,
							EntryClass:         artifactEntryClass,
							FlinkVersion:       artifactFlinkVersion,
							FlinkImageRegistry: artifactFlinkRegistry,
							FlinkImageTag:      artifactFlinkTag,
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
			Expect(string(vpSpec.State)).To(Equal(deploymentState))

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
			Expect(artifact.EntryClass).To(Equal(artifactEntryClass))
			Expect(artifact.MainArgs).To(Equal(artifactArgs))
			Expect(artifact.Kind).To(Equal(artifactKind))
			Expect(artifact.JarUri).To(Equal(artifactJarUri))

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

		Measure("conversion speed", func(b Benchmarker) {
			b.Time("to native", func() {
				_, _ = DeploymentSpecFromNative(vpDeploymentSpec)
			})
		}, 10)
	})
})
