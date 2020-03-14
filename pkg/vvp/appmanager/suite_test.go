package appmanager_test

import (
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t,
		"appmanager",
		[]Reporter{printer.NewlineReporter{}})
}
