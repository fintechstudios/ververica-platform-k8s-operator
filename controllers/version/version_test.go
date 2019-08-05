package version

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Version", func() {
	Describe("GetVersion", func() {
		It("should create a version object", func() {
			version := GetVersion()
			Expect(len(version.BuildDate)).ToNot(BeZero())
		})
	})

	Describe("ToString", func() {
		var version Version
		BeforeEach(func() {
			version = GetVersion()
		})

		It("should create a version string", func() {
			Expect(len(version.ToString())).ToNot(BeZero())
		})
	})
})
