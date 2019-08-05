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

	notFoundErrors := []error{
		DeploymentNotFoundError{Namespace: "grumpy", Name: "goose"},
		errors.New("404 Not Found"),
		apiErrors.NewNotFound(schema.GroupResource{
			Group:    "api.fts.com",
			Resource: "tools",
		}, "hammer"),
	}

	It("should return true for not found errors", func() {
		for _, err := range notFoundErrors {
			Expect(IsNotFoundError(err)).To(BeTrue())
		}
	})
})
