package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Context("operatorVersion", func() {
		Describe("GetVersion", func() {
			It("should create a operatorVersion object", func() {
				version := GetVersion()
				Expect(len(version.BuildDate)).ToNot(BeZero())
			})
		})

		Describe("String", func() {
			var version Version
			BeforeEach(func() {
				version = GetVersion()
			})

			It("should create a operatorVersion string", func() {
				Expect(len(version.String())).ToNot(BeZero())
			})
		})
	})
})
