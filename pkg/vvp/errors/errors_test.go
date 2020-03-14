package vvperrors

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("vvperrors", func() {
	It("should map 403 errors", func() {
		err := errForStatusCode(403, "")
		Expect(errors.Is(err, ErrForbidden)).To(BeTrue())
		Expect(errors.Is(err, ErrBadRequest)).ToNot(BeTrue())
	})

	It("should map 401 errors", func() {
		err := errForStatusCode(401, "")
		Expect(errors.Is(err, ErrUnauthorized)).To(BeTrue())
	})

	It("should return unknown errors for everything else", func() {
		err := errForStatusCode(418, "")
		Expect(errors.Is(err, ErrUnknown)).To(BeTrue())
	})

	It("should include the passed message", func() {
		err := errForStatusCode(418, "teapot")
		Expect(errors.Is(err, ErrUnknown)).To(BeTrue())
		Expect(err.Error()).To(ContainSubstring("teapot"))
	})
})
