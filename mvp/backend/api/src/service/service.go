package service

import (
	"context"
	"errors"
	"log"
	"time"
)

type Service interface {
	Register(context.Context) error
	Webhook(context.Context) error
}

type ServiceStack struct {
	services []Service
}

func New(services ...Service) ServiceStack {
	return ServiceStack{
		services: services,
	}
}

func (ss ServiceStack) Register(ctx context.Context) error {
	errs := make([]error, len(ss.services))
	for _, s := range ss.services {
		err := s.Register(ctx)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return errors.Join(errs...)
	}
	return nil
}

func (ss ServiceStack) Webhook(ctx context.Context) error {
	return errors.New("webhook not available for ServiceStack")
}

func (ss ServiceStack) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(7*24) * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := ss.Register(ctx)
			log.Printf("%+v", err)
		case <-ctx.Done():
			log.Println("Stopping registration due to context cancellation")
		}
	}
}
