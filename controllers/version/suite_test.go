package version

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

func TestVersion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "version")

	RunSpecsWithDefaultAndCustomReporters(t,
		"version",
		[]Reporter{envtest.NewlineReporter{}})

}
