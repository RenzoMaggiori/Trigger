package services

import (
	"context"
	"errors"
	"log"
	"time"
)

type Service interface {
	Register(context.Context) error
}

type ServiceStack struct {
	services []Service
}

func New(services ...Service) Service {
	return ServiceStack{
		services: services,
	}
}

func (ss ServiceStack) Register(ctx context.Context) error {
	ticker := time.NewTicker(time.Duration(7*24) * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, s := range ss.services {
				err := s.Register(ctx)
				if err != nil {
					log.Printf("Error registering: %+v\n", err)
				}
			}
		case <-ctx.Done():
			log.Println("Stopping registration due to context cancellation")
			return errors.New("Stopping registration due to context cancellation")
		}
	}
}
