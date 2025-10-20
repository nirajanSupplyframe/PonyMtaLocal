package user

import (
	"gopro/internal/constants"
	"gopro/internal/service/mail"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetUser(c *gin.Context) {
	res := mail.PostEmailResponseDTO{
		Status:  constants.Success,
		Message: constants.Success,
	}

	c.JSON(200, res)

}
