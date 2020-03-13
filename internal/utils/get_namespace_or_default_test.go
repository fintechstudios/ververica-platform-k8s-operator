package utils

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetNamespaceOrDefault", func() {
	It("should return the default namespace given an empty string", func() {
		Expect(GetNamespaceOrDefault("")).To(Equal(DefaultNamespace))
	})

	It("should return the given namespace when non-empty", func() {
		Expect(GetNamespaceOrDefault("fts")).To(Equal("fts"))
	})
})
