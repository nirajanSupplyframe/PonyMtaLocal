package events

import "time"

type EventType int

const (
	EventQueued EventType = iota
	EventLinked
	EventStatus
)

type Event struct {
	Type      EventType
	RequestID string
	QueueID   string
	MessageID string
	Status    string
	Raw       string
}

type EmailState struct {
	RequestID string    `json:"request_id"`
	QueuedID  string    `json:"queued_id"`
	Status    string    `json:"status"`
	Reason    string    `json:"reason"`
	Updated   time.Time `json:"updated"`
}
