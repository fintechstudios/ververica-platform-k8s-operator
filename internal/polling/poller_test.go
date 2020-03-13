package polling

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Poller", func() {
	var poller *Poller
	BeforeEach(func() {
		poller = NewPoller(func() interface{} {
			return 1
		}, time.Millisecond*1)
	})

	It("should panic if it is started after being stopped", func() {
		poller.Start()
		poller.StopAndBlock()
		Expect(func() {
			poller.Start()
		}).To(Panic())
	})

	It("should close the output channel when stopped", func() {
		poller.Start()
		poller.Stop()
		Eventually(poller.Channel).Should(BeClosed())
		Eventually(poller.IsStopped).Should(BeTrue())
		Eventually(poller.IsFinished).Should(BeTrue())
	})

	It("should close the output channel when stopped and block", func() {
		poller.Start()
		poller.StopAndBlock()
		Expect(poller.Channel).To(BeClosed())
		Expect(poller.IsStopped()).To(BeTrue())
		Expect(poller.IsFinished()).To(BeTrue())
	})

	It("should send output to a channel", func() {
		poller.Start()
		defer poller.Stop()
		Eventually(poller.Channel).Should(Receive(Equal(1)))
	})
})
