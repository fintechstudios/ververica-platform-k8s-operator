/*
Copyright 2019 FinTech Studios, Inc.

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

package v1beta1

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// These tests are written in BDD-style using Ginkgo framework. Refer to
// http://onsi.github.io/ginkgo to learn more.

var _ = Describe("VpDeployment", func() {
	var (
		key              types.NamespacedName
		created, fetched *VpDeployment
	)

	// Add Tests for OpenAPI validation (or additional CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.
	Context("Create API", func() {

		It("should create an object successfully", func() {
			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}
			created = &VpDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: VpDeploymentObjectSpec{
					Metadata: VpMetadata{},
					Spec: VpDeploymentSpec{
						UpgradeStrategy: &VpDeploymentUpgradeStrategy{
							Kind: "STATELESS",
						},
						State: RunningState,
						Template: &VpDeploymentTemplate{
							Spec: &VpDeploymentTemplateSpec{
								Artifact: &VpArtifact{
									Kind:   "JAR",
									JarUri: "https://jars.com/peanut-butter",
								},
							},
						},
					},
					DeploymentTargetName: "dep-target",
				},
			}

			By("creating an API obj")
			Expect(k8sClient.Create(context.TODO(), created)).To(Succeed())

			fetched = &VpDeployment{}
			Expect(k8sClient.Get(context.TODO(), key, fetched)).To(Succeed())
			Expect(fetched).To(Equal(created))

			By("deleting the created object")
			Expect(k8sClient.Delete(context.TODO(), created)).To(Succeed())
			Expect(k8sClient.Get(context.TODO(), key, created)).ToNot(Succeed())
		})

		It("should fail validation when an upgradeStrategy isn't provided", func() {
			created = &VpDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: VpDeploymentObjectSpec{
					Metadata: VpMetadata{},
					Spec: VpDeploymentSpec{
						State: RunningState,
						Template: &VpDeploymentTemplate{
							Spec: &VpDeploymentTemplateSpec{
								Artifact: &VpArtifact{
									Kind:   "JAR",
									JarUri: "https://jars.com/peanut-butter",
								},
							},
						},
					},
				},
			}
			By("creating an API obj")
			err := k8sClient.Create(context.TODO(), created)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("spec.spec.upgradeStrategy"))
		})

		It("should fail validation when a desired state isn't provided", func() {
			created = &VpDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: VpDeploymentObjectSpec{
					Metadata: VpMetadata{},
					Spec: VpDeploymentSpec{
						UpgradeStrategy: &VpDeploymentUpgradeStrategy{
							Kind: "STATELESS",
						},
						Template: &VpDeploymentTemplate{
							Spec: &VpDeploymentTemplateSpec{
								Artifact: &VpArtifact{
									Kind:   "JAR",
									JarUri: "https://jars.com/peanut-butter",
								},
							},
						},
					},
					DeploymentTargetName: "dep-target",
				},
			}
			By("creating an API obj")
			err := k8sClient.Create(context.TODO(), created)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("spec.spec.state"))
		})

		It("should fail validation when an artifact isn't provided", func() {
			created = &VpDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: VpDeploymentObjectSpec{
					Metadata: VpMetadata{},
					Spec: VpDeploymentSpec{
						UpgradeStrategy: &VpDeploymentUpgradeStrategy{
							Kind: "STATELESS",
						},
						State: RunningState,
						Template: &VpDeploymentTemplate{
							Spec: &VpDeploymentTemplateSpec{},
						},
					},
					DeploymentTargetName: "dep-target",
				},
			}
			By("creating an API obj")
			err := k8sClient.Create(context.TODO(), created)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("spec.spec.template.spec.artifact"))
		})
	})
})
