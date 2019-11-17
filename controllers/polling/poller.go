package polling

import (
	"sync"
	"time"
)

type Poller struct {
	Channel      chan<- interface{}
	Poll         func() interface{}
	WaitInterval time.Duration
	isStopped    bool
	isFinished   bool
	group        *sync.WaitGroup
	stopMutex    *sync.Mutex
}

func NewPoller(poll func() interface{}, interval time.Duration) *Poller {
	return &Poller{
		Channel:      make(chan interface{}),
		Poll:         poll,
		WaitInterval: interval,
		isStopped:    false,
		isFinished:   false,
		group:        &sync.WaitGroup{},
		stopMutex:	  &sync.Mutex{},
	}
}

// sendResult forwards on a polling result if the channel is not closed
// which could happen during the polling request
func (p *Poller) sendResult(result interface{}) {
	p.stopMutex.Lock()
	if !p.IsStopped() {
		p.Channel <- result
	}
	p.stopMutex.Unlock()
}

// runPolling is the actual polling mechanism that handles control flow
func (p *Poller) runPolling() {
	for !p.IsStopped() {
		p.sendResult(p.Poll())
		time.Sleep(p.WaitInterval)
	}
	p.group.Done()
	p.isFinished = true
}

// Start starts the polling process -- cannot be restarted after stopping
func (p *Poller) Start() {
	if p.IsStopped() {
		panic("cannot restart poller after it has been stopped")
	}

	p.group.Add(1)
	go p.runPolling()
}

// Stop exits the polling loop on the next attempt, waits for it to finish,
// and closes the channel
func (p *Poller) Stop() {
	p.stopMutex.Lock()
	p.isStopped = true
	close(p.Channel)
	p.stopMutex.Unlock()
}

// StopAndBlock stops the poller and blocks until it is closed
func (p *Poller) StopAndBlock() {
	p.Stop()
	p.group.Wait()
}

// IsFinished returns whether or not the polling worker has completed
func (p *Poller) IsFinished() bool {
	return p.isFinished
}

// IsStopped returns whether or not the polling worker has been stopped
func (p *Poller) IsStopped() bool {
	return p.isStopped
}