package events

import "strings"

func tailLog(sm *StateManager) {

	for {
		logMessage := tailPostfixLog()
		switch {
		case strings.Contains(logMessage, "status= sent"):
			sm.Publish(Event{
				Type:    EventMailDelivered,
				Message: logMessage,
			})
		case strings.Contains(logMessage, "status= bounced"):
			sm.Publish(Event{
				Type:    EventMailBounced,
				Message: logMessage,
			})
		case strings.Contains(logMessage, "status= rejected"):
			sm.Publish(Event{
				Type:    EventMailQueued,
				Message: logMessage,
			})
		}
	}
}

func tailPostfixLog() string {
	return " "
}
