package utils

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AddFinalizer", func() {
	strings := []string{"a", "b"}

	It("should return true when a string is present", func() {
		Expect(ContainsString(strings, "a")).To(BeTrue())
	})

	It("should return false when a string is not present", func() {
		Expect(ContainsString(strings, "c")).To(BeFalse())
	})
})
