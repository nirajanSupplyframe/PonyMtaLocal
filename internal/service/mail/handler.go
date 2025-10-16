package mail

import (
	"gopro/internal/constants"
	"gopro/internal/infra/mail"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	sender mail.Sender
}

func NewHandler() *Handler {
	return &Handler{
		sender: mail.NewPostfixSender(),
	}
}

func (h *Handler) SendMail(c *gin.Context) {

	var req RequestDTO

	if err := c.BindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, getDto(
			constants.Error,
			constants.BadRequest,
		))
		return
	}

	if err := h.sender.SendMail(req.To, req.Subject, req.Body); err != nil {

		c.JSON(500, getDto(
			constants.Error,
			constants.InternalServerError,
		))
	}
	var x = mail.NewParsedMessage()
	c.JSON(http.StatusOK, getDto(
		constants.Success,
		x,
	))
}

func getDto(status, message string) EmailResponseDTO {
	return EmailResponseDTO{
		Status:  status,
		Message: message,
	}
}
