package utils

import (
	"errors"

	appManagerApi "github.com/fintechstudios/ververica-platform-k8s-operator/appmanager-api-client"
	platformApi "github.com/fintechstudios/ververica-platform-k8s-operator/platform-api-client"
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
		appManagerApi.GenericSwaggerError{}.WithStatusCode(404),
		platformApi.GenericSwaggerError{}.WithStatusCode(404),
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
