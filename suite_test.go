package main

import (
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMainPkg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "main")

	RunSpecsWithDefaultAndCustomReporters(t,
		"main",
		[]Reporter{printer.NewlineReporter{}})

}
