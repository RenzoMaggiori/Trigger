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
		ticker:   time.NewTicker(time.Duration(5) * time.Second),
	}
}

func (w TimerWorker) Start() error {
	for {
		select {
		case <-w.ctx.Done():
			return nil
		case <-w.ticker.C:
			err := w.ServicesNotify()
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (w TimerWorker) ServicesNotify() error {
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
	return nil
}

func (w TimerWorker) Stop() {
	w.cancel()
	for _, s := range w.services {
		s.Ticker().Stop()
	}
}
