package utils

import (
	"errors"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

func TestIsNotFoundError(t *testing.T) {
	RegisterFailHandler(Fail)

	Expect(IsNotFoundError(errors.New("server error"))).To(BeFalse())
	Expect(errors.New("404 Not Found")).To(BeTrue())
	Expect(apiErrors.NewNotFound(schema.GroupResource{
		Group: "api.fts.com",
		Resource: "tools",
	}, "hammer")).To(BeTrue())
}
