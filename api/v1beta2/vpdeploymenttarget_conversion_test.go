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
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/annotations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/utils/pointer"

	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These tests are written in BDD-style using Ginkgo framework. Refer to
// http://onsi.github.io/ginkgo to learn more.

var _ = Describe("VpDeploymentTarget conversion", func() {
	jsonPatch := []v1beta1.JSONPatchGeneric{
		{
			Op:    "add",
			Path:  "/hello",
			Value: pointer.StringPtr("gazelle"),
		},
	}
	jsonPatchStr := "[{\"op\":\"add\",\"path\":\"/hello\",\"value\":\"gazelle\"}]"

	It("should convert to the hub", func() {
		// [{"op":"add","path":"/hello","value":"gazelle"}]
		v2 := &VpDeploymentTarget{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "default",
				Annotations: annotations.Create(
					annotations.Pair(annDepTargetPatchSet, jsonPatchStr)),
			},
			Spec: VpDeploymentTargetObjectSpec{
				Metadata: VpMetadata{
					Namespace: "test",
					Annotations: annotations.Create(
						annotations.Pair(annotations.ID, "some-base16-string")),
					Labels: map[string]string{
						"testing": "true",
					},
				},
				Spec: VpDeploymentTargetSpec{
					Kubernetes: VpKubernetesTarget{Namespace: "vvp"},
				},
			},
		}

		v1 := &v1beta1.VpDeploymentTarget{}
		Expect(v2.ConvertTo(v1)).To(Succeed())
		Expect(v1.Spec.Spec.DeploymentPatchSet).ToNot(BeEmpty())
		Expect(v1.Spec.Spec.DeploymentPatchSet).To(BeEquivalentTo(jsonPatch))
	})

	It("should convert from the hub", func() {
		v1 := &v1beta1.VpDeploymentTarget{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "default",
			},
			Spec: v1beta1.VpDeploymentTargetObjectSpec{
				Metadata: v1beta1.VpMetadata{
					Namespace: "test",
					Labels: map[string]string{
						"testing": "true",
					},
					Annotations: annotations.Create(
						annotations.Pair(annotations.ID, "some-base16-string"),
					),
				},
				Spec: v1beta1.VpDeploymentTargetSpec{
					Kubernetes:         v1beta1.VpKubernetesTarget{Namespace: "vvp"},
					DeploymentPatchSet: jsonPatch,
				},
			},
		}

		v2 := &VpDeploymentTarget{}
		Expect(v2.ConvertFrom(v1)).To(Succeed())
		Expect(annotations.Get(v2.Annotations, annDepTargetPatchSet)).To(Equal(jsonPatchStr))
	})

	It("should covert to and from the hub without losing information", func() {
		v2 := &VpDeploymentTarget{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "default",
				Annotations: annotations.Create(
					annotations.Pair(annDepTargetPatchSet, jsonPatchStr)),
			},
			Spec: VpDeploymentTargetObjectSpec{
				Metadata: VpMetadata{
					Namespace: "test",
					Annotations: annotations.Create(
						annotations.Pair(annotations.ID, "some-base16-string")),
					Labels: map[string]string{
						"testing": "true",
					},
				},
				Spec: VpDeploymentTargetSpec{
					Kubernetes: VpKubernetesTarget{Namespace: "vvp"},
				},
			},
		}
		v1 := &v1beta1.VpDeploymentTarget{}
		Expect(v2.ConvertTo(v1)).To(Succeed())
		v2Clone := &VpDeploymentTarget{}
		Expect(v2Clone.ConvertFrom(v1)).To(Succeed())
		Expect(v2).To(BeEquivalentTo(v2Clone))
	})
})
