package converters

import (
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConverters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t,
		"converters",
		[]Reporter{envtest.NewlineReporter{}})
}
