package events

import (
	"fmt"
	"time"
)

func (sm *StateManager) eventLoop() {

	byRequest := map[string]*EmailState{}
	queueToReq := map[string]string{}
	msdIdToReq := map[string]string{}

	for {
		select {
		case event := <-sm.eventCh:
			switch event.Type {
			case EventQueued:
				if st, ok := byRequest[event.RequestID]; ok {
					st.QueuedID = event.QueueID
					st.Updated = time.Now()
				} else {
					byRequest[event.RequestID] = &EmailState{
						RequestID: event.RequestID,
						QueuedID:  event.QueueID,
						Status:    "queued",
						Updated:   time.Now(),
					}
				}
				//Map message id to request
				msdIdToReq[event.RequestID] = event.RequestID
				println("email QUEUED with request id :", event.RequestID)

			case EventLinked:
				println("email LINKED event loop with request id :", event.RequestID)
				if event.RequestID == "" || event.QueueID == "" {
					break
				}
				queueToReq[event.QueueID] = event.RequestID
				if st, ok := byRequest[event.RequestID]; ok {
					st.QueuedID = event.QueueID
					st.Updated = time.Now()
				} else {
					byRequest[event.RequestID] = &EmailState{
						RequestID: event.RequestID,
						QueuedID:  event.QueueID,
						Status:    "linked and queued",
						Updated:   time.Now(),
					}
				}

			case EventStatus:
				q := event.QueueID
				reqID, ok := queueToReq[q]
				if !ok {
					reqID = event.RequestID
					if reqID == "" {
						//orphan status
						break
					}
				}
				st, exists := byRequest[reqID]
				if !exists {
					st = &EmailState{
						RequestID: reqID,
						QueuedID:  q,
					}
					byRequest[reqID] = st

				}
				st.QueuedID = q
				st.Status = event.Status
				st.Reason = event.Raw
				st.Updated = time.Now()
				byRequest[reqID] = st
			}

		case req := <-sm.requestCh:
			if st, ok := byRequest[req.requestID]; ok {
				copied := *st
				req.resp <- &copied
			} else {
				req.resp <- nil
			}

		case <-sm.doneCh:
			fmt.Println("Event loop shutting down...")
			return
		}
	}
}
