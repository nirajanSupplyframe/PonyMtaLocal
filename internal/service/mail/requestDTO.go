package mail

type RequestDTO struct {
	To      string `json:"To"`
	Subject string `json:"Subject"`
	Body    string `json:"Body"`
}
