package utils

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

func TestRemoveString(t *testing.T) {
	RegisterFailHandler(Fail)

	strings := []string{"a", "b"}
	Expect(len(RemoveString(strings, "c"))).To(Equal(2))
	Expect(len(RemoveString(strings, "a"))).To(Equal(1))
}
