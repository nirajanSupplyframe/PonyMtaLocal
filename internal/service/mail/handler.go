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
		return
	}

	var x = mail.NewParsedMessage()
	if x == nil {
		c.JSON(500, getDto(
			constants.Error,
			constants.InternalServerError,
		))
		return
	}
	if x.DSN[0] == '2' {
		c.JSON(http.StatusOK, getDto(
			constants.Success,
			x.Message,
		))
	} else if x.DSN[0] == '4' {
		c.JSON(http.StatusBadRequest, getDto(
			constants.Error,
			x.Message,
		))
	} else if x.DSN[0] == '5' {
		c.JSON(http.StatusInternalServerError, getDto(
			constants.Error,
			x.Message,
		))
	}
}

func getDto(status, message string) EmailResponseDTO {
	return EmailResponseDTO{
		Status:  status,
		Message: message,
	}
}
