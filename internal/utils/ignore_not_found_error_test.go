package utils

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var _ = Describe("IgnoreNotFoundError", func() {
	It("should return nil for not found errors", func() {
		Expect(IgnoreNotFoundError(DeploymentNotFoundError{Namespace: "rotten", Name: "eggs"})).To(BeNil())
	})

	It("should return the error for all others", func() {
		err := apiErrors.NewServerTimeout(schema.GroupResource{
			Group:    "api.fts.com",
			Resource: "tools",
		},
			"create",
			5)

		Expect(IgnoreNotFoundError(err)).To(Equal(err))
	})
})
