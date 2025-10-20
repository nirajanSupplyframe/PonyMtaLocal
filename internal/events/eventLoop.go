package events

import "fmt"

func (sm *StateManager) eventLoop() {

	for {
		select {
		case event := <-sm.eventCh:
			sm.handle(event)
		case <-sm.doneCh:
			fmt.Println("Event loop shutting down...")
			return
		}
	}
}

func (sm *StateManager) handle(event Event) {
	switch event.Type {
	case EventMailDelivered:
		fmt.Println("Email Delivered : ", event.Message)
	case EventMailQueued:
		fmt.Println("Email in Queue : ", event.Message)
	case EventMailBounced:
		fmt.Println("Email Bounced : ", event.Message)
	}
}
