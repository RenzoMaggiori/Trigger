package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"trigger.com/api/src/endpoints/gmail"
	"trigger.com/api/src/middleware"
	"trigger.com/api/src/router"
	"trigger.com/api/src/services"
)

type server struct {
	wrapper *http.Server
	ctx     context.Context
}

func Create(port int64, ctx context.Context) (*server, error) {
	middleware := middleware.Create(
		middleware.Cors,
	)
	router, err := router.Create(ctx)
	if err != nil {
		return nil, err
	}

	return &server{
		wrapper: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: middleware(router),
		},
		ctx: ctx,
	}, nil
}

func (s *server) Start() {
	go services.New(
		gmail.Model{},
	).Register(s.ctx)

	log.Printf("Listening on %s\n", s.wrapper.Addr)
	if err := s.wrapper.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", s.wrapper.Addr, err)
	}
}

func (s *server) Stop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down server...")
	if err := s.wrapper.Close(); err != nil {
		log.Fatal("Server shutdown failed:", err)
	}
	log.Println("Server gracefully stopped")
}
