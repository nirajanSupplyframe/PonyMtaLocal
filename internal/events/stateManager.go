package events

type StateManager struct {
	eventCh chan Event
	doneCh  chan struct{}
}

func NewStateManager() *StateManager {
	return &StateManager{
		eventCh: make(chan Event),
		doneCh:  make(chan struct{}),
	}
}

func (sm *StateManager) Start() {
	go sm.eventLoop()

}

func (sm *StateManager) Stop() {
	close(sm.doneCh)
}

func (sm *StateManager) Publish(event Event) {
	sm.eventCh <- event
}
