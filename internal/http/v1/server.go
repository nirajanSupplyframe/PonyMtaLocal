package v1

import (
	"gopro/internal/config"
	"gopro/internal/service/mail"
	"gopro/internal/service/user"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	cfg    config.Config
}

func NewServer(cfg config.Config) *Server {

	r := gin.Default()

	userHandler := user.NewHandler()
	mailHandler := mail.NewHandler()

	r.GET("/api/v1/user", userHandler.GetUser)
	r.POST("/api/v1/mail", mailHandler.SendMail)

	return &Server{engine: r, cfg: cfg}

}

func (s *Server) Start() {
	err := s.engine.Run(":" + s.cfg.Port)
	if err != nil {
		return
	}
}
