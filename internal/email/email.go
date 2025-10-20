package email

type ProcessingState string

const (
	ProcessingStateEnqueueing ProcessingState = "enqueueing"
	ProcessingStatSending     ProcessingState = "sending"
	ProcessingStatSent        ProcessingState = "sent"
	ProcessingStatFailed      ProcessingState = "failed"
)

type MtaStatus string

const (
	MtaStatusUnknown  MtaStatus = "unknown"
	MtaStatusSent     MtaStatus = "sent"
	MtaStatusDeferred MtaStatus = "deferred"
	MtaStatusBounced  MtaStatus = "bounced"
	MtaStatusExpired  MtaStatus = "expired"
)

type StatusEmailObject struct {
	Id           int
	State        ProcessingState
	PostfixId    string
	Status       MtaStatus
	StatusTextID int
	DomainId     int
	Created      string
	Updated      string
	Log          string
}
