package v1

import (
	"gopro/internal/events"
	mail2 "gopro/internal/infra/mail"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	sm     *events.StateManager
	sender mail2.Sender
}

func NewServer(sm *events.StateManager, s mail2.Sender) *Server {
	return &Server{sm: sm, sender: s}
}

type sendReq struct {
	To      string `json:"to" binding:"required,email" `
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

type sendResp struct {
	RequestID string `json:"requestID"`
	Status    string `json:"status"`
}

func (s *Server) RegisterRoutes(r *gin.Engine) {
	r.POST("/send", s.handleSend)
	r.GET("/status/:id", s.handleStatus)
}

func (s *Server) handleSend(c *gin.Context) {
	var req sendReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Sends mail and injects message-id to mail from body param.
	reqID, err := s.sender.SendMail(req.To, req.Subject, req.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//notify state manager about this request, and that it exists.
	println("email QUEUED (handle send before publish) with request id :", reqID)
	s.sm.Publish(events.Event{
		Type:      events.EventQueued,
		RequestID: reqID,
		Status:    "queued",
	})

	c.JSON(http.StatusAccepted, sendResp{RequestID: reqID, Status: s.sm.GetStatus(reqID).Status})

}

func (s *Server) handleStatus(c *gin.Context) {
	id := c.Param("id")
	st := s.sm.GetStatus(id)
	if st == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not found"})
		return
	}
	c.JSON(http.StatusOK, st)
}
