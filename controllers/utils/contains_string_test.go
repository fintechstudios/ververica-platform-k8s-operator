package utils

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

func TestContainsString(t *testing.T) {
	RegisterFailHandler(Fail)

	strings := []string{"a", "b"}
	Expect(ContainsString(strings, "c")).To(BeFalse())
	Expect(ContainsString(strings, "a")).To(BeTrue())
}
