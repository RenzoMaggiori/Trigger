package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"trigger.com/api/src/router"
)

type Server struct {
	wrapper *http.Server
}

func Create(port int64) *Server {
	router := router.Create()
	return &Server{
		wrapper: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
	}
}

func (s *Server) Start() {
	if err := s.wrapper.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", s.wrapper.Addr, err)
	}
}

func (s *Server) Stop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down server...")
	if err := s.wrapper.Close(); err != nil {
		log.Fatal("Server shutdown failed:", err)
	}
	log.Println("Server gracefully stopped")
}
