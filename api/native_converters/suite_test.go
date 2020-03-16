package native_converters_test

import (
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConverters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t,
		"Native Converters",
		[]Reporter{printer.NewlineReporter{}})
}
