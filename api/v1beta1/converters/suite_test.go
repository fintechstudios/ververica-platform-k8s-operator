package converters

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConverters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "converters")
}
