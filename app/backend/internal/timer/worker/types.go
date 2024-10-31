package worker

import (
	"context"
	"time"
)

type Worker interface {
	// Careful this function is blocking
	Start() error
	Stop()
}

type TimerService interface {
	Notify(context.Context) error
	Ticker() *time.Ticker
}

type TimerWorker struct {
	services []TimerService
	ctx      context.Context
	cancel   context.CancelFunc
	ticker   *time.Ticker
}

type MinuteService struct {
	ticker *time.Ticker
}

type HourService struct {
	ticker *time.Ticker
}

type DayService struct {
	ticker *time.Ticker
}
