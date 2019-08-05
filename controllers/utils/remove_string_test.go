package utils

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var _ = Describe("RemoveString", func() {
	var strings []string

	BeforeEach(func() {
		strings = []string{"a", "b"}
	})

	It("should remove a string if it is included", func() {
		Expect(len(RemoveString(strings, "c"))).To(Equal(2))
	})

	It("should not remove anything if a string is not included", func() {
		Expect(len(RemoveString(strings, "a"))).To(Equal(1))
	})
})
