package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"trigger.com/trigger/pkg/middleware"
)

type Server struct {
	wrapper *http.Server
}

func Create(
	router *http.ServeMux,
	middlewares middleware.Middleware,
	port int64,
) (*Server, error) {

	return &Server{
		wrapper: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: middlewares(router),
		},
	}, nil
}

func (s *Server) Start() {
	fmt.Printf("Listening on http://localhost%s\n", s.wrapper.Addr)
	if err := s.wrapper.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on http://localhost%s: %v\n", s.wrapper.Addr, err)
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
