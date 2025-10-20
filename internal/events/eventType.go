package events

type EventType int

const (
	EventMailQueued EventType = iota
	EventMailDelivered
	EventMailBounced
)

type Event struct {
	Type    EventType
	Message any
}
