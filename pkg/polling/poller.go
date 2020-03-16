package polling

import (
	"sync"
	"time"
)

type commandResult struct{}

// FinishedResult is a sentinel that can be returned by a PollerFunc to signal polling should exit
var FinishedResult = &commandResult{}

// PollerFunc is a function to be polled
type PollerFunc func() interface{}

// Poller represents everything needed for polling a function
type Poller struct {
	Channel      chan interface{}
	Poll         PollerFunc
	WaitInterval time.Duration
	status       status
	group        *sync.WaitGroup
	statusMutex  *sync.Mutex
}

type status = string

const (
	runnable = status("runnable") // 1
	running  = status("running")  // 2
	stopped  = status("stopped")  // 3
	finished = status("finished") // 4
)

// NewPoller creates a new poller for a function with a polling interval
// NOTE: the polling function cannot return `nil`
func NewPoller(poll PollerFunc, interval time.Duration) *Poller {
	return &Poller{
		Channel:      make(chan interface{}),
		Poll:         poll,
		WaitInterval: interval,
		status:       runnable,
		group:        &sync.WaitGroup{},
		statusMutex:  &sync.Mutex{},
	}
}

// sendResult forwards on a polling result if the channel is not closed
// which could happen during the polling request
func (p *Poller) sendResult(result interface{}) {
	p.statusMutex.Lock()
	defer p.statusMutex.Unlock()
	if !p.IsStopped() {
		p.Channel <- result
	}
}

// runPolling is the actual polling mechanism that handles control flow
func (p *Poller) runPolling() {
	for !p.IsDone() {
		if res := p.Poll(); res != nil {
			if cmdRes, ok := res.(commandResult); ok && &cmdRes == FinishedResult {
				p.Stop()
				break
			}

			p.sendResult(res)
		}
		time.Sleep(p.WaitInterval)
	}
	p.group.Done()
	p.statusMutex.Lock()
	defer p.statusMutex.Unlock()
	p.status = finished
}

// Start starts the polling process -- cannot be restarted after stopping
func (p *Poller) Start() {
	// TODO: make this less panic-driven by accommodating this case
	if p.IsDone() {
		panic("cannot restart poller after it has been stopped")
	}

	if p.IsRunning() {
		// already running
		return
	}
	p.statusMutex.Lock()
	defer p.statusMutex.Unlock()
	if p.IsRunning() {
		// changed to running by another routine
		return
	}
	p.status = running

	p.group.Add(1)
	go p.runPolling()
}

// Stop exits the polling loop on the next attempt, waits for it to finish,
// and closes the channel
func (p *Poller) Stop() {
	p.statusMutex.Lock()
	defer p.statusMutex.Unlock()

	if p.IsStopped() {
		// already been closed
		return
	}

	p.status = stopped
	close(p.Channel)
}

// StopAndBlock stops the poller and blocks until it is closed
func (p *Poller) StopAndBlock() {
	p.Stop()
	p.group.Wait()
}

// IsRunning returns whether or not the poller is able to be started
func (p *Poller) IsRunning() bool {
	return p.status == running
}

// IsFinished returns whether or not the polling worker has completed
func (p *Poller) IsFinished() bool {
	return p.status == finished
}

// IsStopped returns whether or not the polling worker has been stoped
func (p *Poller) IsStopped() bool {
	return p.status == stopped
}

// IsDone returns whether or not the polling worker has been stopped or is finished
func (p *Poller) IsDone() bool {
	return p.status == stopped || p.IsFinished()
}
