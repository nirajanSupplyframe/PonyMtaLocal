package main

import (
	"gopro/internal/events"
	v1 "gopro/internal/http/v1"
	mail2 "gopro/internal/infra/mail"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//Creates a new event events manager
	sm := events.NewStateManager()
	sm.Start()

	//Start the log tailer (parses Postfix log and publishes it to the event loop)
	go events.TailLog(sm)

	s := mail2.NewPostfixSender("localhost:25")

	//Create gin engine and register routes
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	server := v1.NewServer(sm, s)
	server.RegisterRoutes(r)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  1000000 * time.Second,
		WriteTimeout: 1000000 * time.Second,
	}
	log.Println("listening on :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v ", err)
	}
}
