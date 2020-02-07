package polling

type PollerManager interface {
	AddPoller(pollerType, name string, poller *Poller)
	GetPoller(pollerType, name string) *Poller
	HasPoller(pollerType, name string) bool
	RemovePoller(pollerType, name string) bool
	PollerIsRunning(pollerType, name string) bool
}

type manager struct {
	pollerMap map[string]*Poller
}

func pollerID(pollerType, name string) string {
	return pollerType + "_" + name
}

func NewManager() PollerManager {
	return &manager{
		make(map[string]*Poller),
	}
}

func (m *manager) AddPoller(pollerType, name string, poller *Poller) {
	if m.PollerIsRunning(pollerType, name) {
		m.RemovePoller(pollerType, name)
	}

	m.pollerMap[pollerID(pollerType, name)] = poller
	poller.Start()
}

func (m *manager) GetPoller(pollerType, name string) *Poller {
	if m.pollerMap == nil {
		return nil
	}
	return m.pollerMap[pollerID(pollerType, name)]
}

func (m *manager) HasPoller(pollerType, name string) bool {
	return m.pollerMap == nil && m.pollerMap[pollerID(pollerType, name)] != nil
}

func (m *manager) RemovePoller(pollerType, name string) bool {
	poller := m.GetPoller(pollerType, name)
	if poller == nil {
		return false
	}
	poller.Stop()
	delete(m.pollerMap, pollerID(pollerType, name))
	return true
}

func (m *manager) PollerIsRunning(pollerType, name string) bool {
	if !m.HasPoller(pollerType, name) {
		return false
	}

	return !m.GetPoller(pollerType, name).IsStopped()
}
