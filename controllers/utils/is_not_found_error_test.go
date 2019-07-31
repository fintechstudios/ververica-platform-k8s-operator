package utils

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var _ = Describe("IsNotFoundError", func() {

	It("should return false for non 404 errors", func() {
		Expect(IsNotFoundError(errors.New("server error"))).To(BeFalse())
	})

	It("should return true for 404 Swagger Errors", func() {
		Expect(IsNotFoundError(errors.New("404 Not Found"))).To(BeTrue())
	})

	It("should return true for K8s api errors", func() {
		Expect(IsNotFoundError(apiErrors.NewNotFound(schema.GroupResource{
			Group:    "api.fts.com",
			Resource: "tools",
		}, "hammer"))).To(BeTrue())
	})
})
