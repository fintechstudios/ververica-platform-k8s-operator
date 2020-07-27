/*
Copyright 2020 FinTech Studios, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta2

import (
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/annotations"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	"time"
)

// These tests are written in BDD-style using Ginkgo framework. Refer to
// http://onsi.github.io/ginkgo to learn more.

var _ = Describe("VpDeployment conversion", func() {
	It("should convert to the hub", func() {
		v2 := &VpDeployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "default",
			},
			Spec: VpDeploymentObjectSpec{
				Metadata: VpMetadata{
					Namespace: "testing",
					Annotations: annotations.Create(
						annotations.Pair(annotations.ID, "some-base16-string")),
					Labels: map[string]string{
						"testing": "true",
					},
				},
				Spec: VpDeploymentSpec{
					UpgradeStrategy: &VpDeploymentUpgradeStrategy{
						Kind: "STATELESS",
					},
					State: RunningState,
					Template: &VpDeploymentTemplate{
						Spec: &VpDeploymentTemplateSpec{
							Artifact: &VpArtifact{
								Kind:   "JAR",
								JarURI: "https://jars.com/peanut-butter",
							},
							Kubernetes: &VpKubernetesOptions{
								Pods: &VpPodSpec{
									EnvVars: []core.EnvVar{
										{
											Name:  "TEST_ENV",
											Value: "TEST_VALUE",
										},
									},
								},
							},
						},
					},
				},
				DeploymentTargetName: "dep-target",
			},
			Status: &VpDeploymentStatus{
				State:   "",
				Running: nil,
			},
		}

		v1 := &v1beta1.VpDeployment{}
		Expect(v2.ConvertTo(v1)).To(Succeed())
	})

	It("should convert from the hub", func() {
		v1 := &v1beta1.VpDeployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "default",
			},
			Spec: v1beta1.VpDeploymentObjectSpec{
				Metadata: v1beta1.VpMetadata{},
				Spec: v1beta1.VpDeploymentSpec{
					UpgradeStrategy: &v1beta1.VpDeploymentUpgradeStrategy{
						Kind: "STATELESS",
					},
					StartFromSavepoint: &v1beta1.VpDeploymentStartFromSavepoint{Kind: "something-feels-wrong"},
					State:              v1beta1.RunningState,
					Template: &v1beta1.VpDeploymentTemplate{
						Spec: &v1beta1.VpDeploymentTemplateSpec{
							Artifact: &v1beta1.VpArtifact{
								Kind:   "JAR",
								JarURI: "https://jars.com/peanut-butter",
							},
							Kubernetes: &v1beta1.VpKubernetesOptions{
								Pods: &v1beta1.VpPodSpec{
									EnvVars: []core.EnvVar{
										{
											Name:  "TEST_ENV",
											Value: "TEST_VALUE",
										},
									},
								},
							},
						},
					},
				},
				DeploymentTargetName: "dep-target",
			},
			Status: v1beta1.VpDeploymentStatus{State: v1beta1.RunningState},
		}

		v2 := &VpDeployment{}
		Expect(v2.ConvertFrom(v1)).To(Succeed())
		Expect(annotations.Has(v2.Annotations, annDepStartFromSavepoint)).To(BeTrue())
	})

	It("should covert to and from the hub without losing information", func() {
		v2 := &VpDeployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "foo",
				Namespace:   "default",
				Annotations: annotations.Create(),
			},
			Spec: VpDeploymentObjectSpec{
				DeploymentTargetName: "dep-target",
				Metadata: VpMetadata{
					Namespace: "testing-b",
					Annotations: annotations.Create(
						annotations.Pair("testing", "true")),
					Labels: map[string]string{
						"license/testing": "true",
					},
				},
				Spec: VpDeploymentSpec{
					UpgradeStrategy: &VpDeploymentUpgradeStrategy{
						Kind: "STATELESS",
					},
					RestoreStrategy: &VpDeploymentRestoreStrategy{
						Kind:                  "NONE",
						AllowNonRestoredState: true,
					},
					State:                        RunningState,
					DeploymentTargetID:           "an-id",
					MaxSavepointCreationAttempts: pointer.Int32Ptr(4),
					MaxJobCreationAttempts:       pointer.Int32Ptr(3),
					Template: &VpDeploymentTemplate{
						Metadata: &VpDeploymentTemplateMetadata{
							Annotations: annotations.Create(
								annotations.Pair(annotations.ResourceVersion, "43"),
							),
						},
						Spec: &VpDeploymentTemplateSpec{
							Artifact: &VpArtifact{
								Kind:                 "JAR",
								JarURI:               "https://jars.com/peanut-butter",
								MainArgs:             "--help",
								EntryClass:           "com.fintechstudios.Flubber",
								FlinkVersion:         "1.10",
								FlinkImageTag:        "1.10-stream2",
								FlinkImageRepository: "somewhere",
								FlinkImageRegistry:   "ververica.com/",
							},
							Parallelism:          pointer.Int32Ptr(2),
							NumberOfTaskManagers: pointer.Int32Ptr(4),
							Resources: map[string]VpResourceSpec{
								"tm": {
									CPU:    resource.MustParse("4"),
									Memory: pointer.StringPtr("4GB"),
								},
								"jm": {
									CPU:    resource.MustParse("2"),
									Memory: pointer.StringPtr("2GB")},
							},
							FlinkConfiguration: map[string]string{
								"high-availability.storageDir": "s3://flink/haState",
							},
							Logging: &VpLogging{
								Log4jLoggers: map[string]string{
									"com.fintechstudios": "DEBUG",
								},
							},
							Kubernetes: &VpKubernetesOptions{
								Pods: &VpPodSpec{
									Annotations: annotations.Create(
										annotations.Pair(annotations.DeploymentID, "some-id")),
									EnvVars: []core.EnvVar{
										{
											Name:  "TEST_DATA",
											Value: "BORING",
										},
										{
											Name: "API_KEY",
											ValueFrom: &core.EnvVarSource{
												SecretKeyRef: &core.SecretKeySelector{
													LocalObjectReference: core.LocalObjectReference{
														Name: "some-secret",
													},
													Key:      "api-key",
													Optional: pointer.BoolPtr(true),
												},
											},
										},
									},
									Labels: map[string]string{
										"k8s.io/part-of": "fraud-detection-system",
									},
									VolumeMounts: []VpVolumeAndMount{
										{
											Name: "a-mount",
											Volume: &core.Volume{
												Name: "a-volume",
												VolumeSource: core.VolumeSource{
													ConfigMap: &core.ConfigMapVolumeSource{
														LocalObjectReference: core.LocalObjectReference{
															Name: "some-config",
														},
													},
												},
											},
											VolumeMount: &core.VolumeMount{
												Name:      "a-volume-mount",
												ReadOnly:  true,
												MountPath: "/run/config",
												SubPath:   "some-config",
											},
										},
									},
									NodeSelector: map[string]string{
										"disktype": "ssd",
									},
									SecurityContext: &core.PodSecurityContext{
										RunAsUser:    pointer.Int64Ptr(1000),
										RunAsGroup:   pointer.Int64Ptr(1000),
										RunAsNonRoot: pointer.BoolPtr(true),
									},
									ImagePullSecrets: []core.LocalObjectReference{
										{Name: "an-image-pull-secret"},
									},
									Affinity: &core.Affinity{
										NodeAffinity: &core.NodeAffinity{
											RequiredDuringSchedulingIgnoredDuringExecution: &core.NodeSelector{
												NodeSelectorTerms: []core.NodeSelectorTerm{
													{
														MatchExpressions: []core.NodeSelectorRequirement{
															{
																Key:      "kubernetes.io/e2e-az-name",
																Operator: core.NodeSelectorOpIn,
																Values:   []string{"e2e-az1"},
															},
														},
													},
												},
											},
											PreferredDuringSchedulingIgnoredDuringExecution: nil,
										},
										PodAffinity: &core.PodAffinity{
											RequiredDuringSchedulingIgnoredDuringExecution: nil,
											PreferredDuringSchedulingIgnoredDuringExecution: []core.WeightedPodAffinityTerm{
												{
													Weight: 3,
													PodAffinityTerm: core.PodAffinityTerm{
														LabelSelector: &metav1.LabelSelector{
															MatchLabels: map[string]string{
																"k8s.io/pod": "friendly",
															},
														},
														Namespaces:  []string{"kube-system", "default"},
														TopologyKey: "failure-domain.beta.kubernetes.io/zone",
													},
												},
											},
										},
										PodAntiAffinity: &core.PodAntiAffinity{
											RequiredDuringSchedulingIgnoredDuringExecution: []core.PodAffinityTerm{
												{
													LabelSelector: &metav1.LabelSelector{
														MatchLabels: map[string]string{
															"k8s.io/pod": "not-friendly",
														},
													},
													Namespaces:  []string{"kube-system", "default"},
													TopologyKey: "failure-domain.beta.kubernetes.io/zone",
												},
											},
										},
									},
									Tolerations: []core.Toleration{
										{
											Key:      "example-key",
											Operator: core.TolerationOpExists,
											Effect:   "NoSchedule",
										},
									},
								},
							},
						},
					},
				},
			},
			Status: &VpDeploymentStatus{
				State: FinishedState,
				Running: &VpDeploymentRunningStatus{
					Conditions: []VpDeploymentRunningCondition{
						{
							Type:    "ClusterUnreachable",
							Message: "Unknown cluster failure",
							Status:  "Unknown",
							Reason:  "Failed to contact cluster",
							LastTransitionTime: metav1.NewTime(
								utils.MustParseTime(time.RFC3339, "2020-01-02T15:04:05Z"),
							),
							LastUpdateTime: metav1.NewTime(
								utils.MustParseTime(time.RFC3339, "2020-01-02T14:00:00Z"),
							),
						},
						{
							Type:    "JobFailing",
							Message: "Something is not happy",
							Status:  "True",
							Reason:  "error message here?",
							LastTransitionTime: metav1.NewTime(
								utils.MustParseTime(time.RFC3339, "2021-01-02T15:04:05Z"),
							),
							LastUpdateTime: metav1.NewTime(
								utils.MustParseTime(time.RFC3339, "2021-01-02T14:00:00Z"),
							),
						},
					},
					JobID: "a-job-id",
					TransitionTime: metav1.NewTime(
						utils.MustParseTime(time.RFC3339, "2020-01-02T15:04:05Z"),
					),
				},
			},
		}
		v1 := &v1beta1.VpDeployment{}
		Expect(v2.ConvertTo(v1)).To(Succeed())
		v2Clone := &VpDeployment{}
		Expect(v2Clone.ConvertFrom(v1)).To(Succeed())
		Expect(v2.ObjectMeta).To(BeEquivalentTo(v2Clone.ObjectMeta))
		Expect(v2.Spec.DeploymentTargetName).To(BeEquivalentTo(v2Clone.Spec.DeploymentTargetName))
		Expect(v2.Spec.Metadata).To(BeEquivalentTo(v2Clone.Spec.Metadata))
		Expect(v2.Spec.Spec.State).To(BeEquivalentTo(v2Clone.Spec.Spec.State))
		Expect(v2.Spec.Spec.UpgradeStrategy).To(BeEquivalentTo(v2Clone.Spec.Spec.UpgradeStrategy))
		Expect(v2.Spec.Spec.RestoreStrategy).To(BeEquivalentTo(v2Clone.Spec.Spec.RestoreStrategy))
		Expect(v2.Spec.Spec.DeploymentTargetID).To(BeEquivalentTo(v2Clone.Spec.Spec.DeploymentTargetID))
		Expect(v2.Spec.Spec.MaxSavepointCreationAttempts).To(BeEquivalentTo(v2Clone.Spec.Spec.MaxSavepointCreationAttempts))
		Expect(v2.Spec.Spec.MaxJobCreationAttempts).To(BeEquivalentTo(v2Clone.Spec.Spec.MaxJobCreationAttempts))
		Expect(v2.Spec.Spec.Template.Metadata).To(BeEquivalentTo(v2Clone.Spec.Spec.Template.Metadata))
		Expect(v2.Spec.Spec.Template.Spec).To(BeEquivalentTo(v2Clone.Spec.Spec.Template.Spec))
		k8sSpec := v2.Spec.Spec.Template.Spec.Kubernetes
		Expect(k8sSpec.Pods.EnvVars).To(BeEquivalentTo(v2Clone.Spec.Spec.Template.Spec.Kubernetes.Pods.EnvVars))
	})
})
