package events

type StateManager struct {
	eventCh   chan Event
	requestCh chan statusRequest
	doneCh    chan struct{}
}

type statusRequest struct {
	requestID string
	resp      chan *EmailState
}

func NewStateManager() *StateManager {
	return &StateManager{
		eventCh:   make(chan Event),
		requestCh: make(chan statusRequest),
		doneCh:    make(chan struct{}),
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

func (sm *StateManager) GetStatus(requestID string) *EmailState {
	req := statusRequest{
		requestID: requestID,
		resp:      make(chan *EmailState),
	}
	sm.requestCh <- req
	return <-req.resp
}
