package utils

import (
	"errors"
	vvperrors "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var _ = Describe("IsNotFoundError", func() {

	It("should return false for non 404 errors", func() {
		Expect(IsNotFoundError(errors.New("server error"))).To(BeFalse())
	})

	notFoundErrors := []error{
		vvperrors.ErrNotFound,
		k8serrors.NewNotFound(schema.GroupResource{
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
