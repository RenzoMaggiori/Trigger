package worker

import (
	"errors"
	"log"
	"time"

	"golang.org/x/net/context"
)

var (
	errChannelClosed = errors.New("channel closed unexpectedly")
)

func Create(ctx context.Context, cancel context.CancelFunc) Worker {
	services := []TimerService{
		MinuteNotifier(),
		HourNotifier(),
		DayNotifier(),
	}
	return TimerWorker{
		services: services,
		ctx:      ctx,
		cancel:   cancel,
	}
}

func (w TimerWorker) Start() error {
	for {
		select {
		case <-w.ctx.Done():
			return nil
		default:
			for _, s := range w.services {
				select {
				case _, ok := <-s.Ticker().C:
					if !ok {
						return errChannelClosed
					}
					log.Printf("Notifying at %s\n", time.Now().Format(time.RFC850))
					go s.Notify(w.ctx)
				default:
				}
			}
		}
	}
}

func (w TimerWorker) Stop() {
	w.cancel()
	for _, s := range w.services {
		s.Ticker().Stop()
	}
}
