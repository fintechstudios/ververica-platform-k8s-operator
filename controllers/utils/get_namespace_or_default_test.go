package utils

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

func TestGetNamespaceOrDefault(t *testing.T) {
	RegisterFailHandler(Fail)

	Expect(GetNamespaceOrDefault("")).To(Equal(DefaultNamespace))
	Expect(GetNamespaceOrDefault("fts")).To(Equal("fts"))
}
