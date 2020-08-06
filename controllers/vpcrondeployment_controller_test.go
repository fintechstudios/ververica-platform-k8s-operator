package controllers

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("VpCronDeployment Controller", func() {
	var reconciler VpCronDeploymentReconciler

	BeforeEach(func() {
		reconciler = VpCronDeploymentReconciler{
			Client:           k8sClient,
			Log:              logger,
		}
	})

	It("should update a k8s deployment target with a VP deployment target", func() {

	})
})
