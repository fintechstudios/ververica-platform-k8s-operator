package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

func TestMainPkg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "main")

	RunSpecsWithDefaultAndCustomReporters(t,
		"main",
		[]Reporter{envtest.NewlineReporter{}})

}
