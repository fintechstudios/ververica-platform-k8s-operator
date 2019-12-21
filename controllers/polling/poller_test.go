package polling

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Poller", func() {
	It("should panic if it is started after being stopped", func() {
		poller := NewPoller(func() interface{} {
			return 1
		}, time.Millisecond*1)
		poller.Start()
		poller.StopAndBlock()
		Expect(func() {
			poller.Start()
		}).To(Panic())
	})

	It("should close the output channel when stopped", func() {
		poller := NewPoller(func() interface{} {
			return 1
		}, time.Millisecond*1)
		poller.Start()
		poller.StopAndBlock()
		Expect(poller.IsStopped()).To(BeTrue())
		Expect(poller.IsFinished()).To(BeTrue())
	})
})
