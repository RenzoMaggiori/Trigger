package trigger

import (
	"time"

	"trigger.com/trigger/pkg/action"
)

type TimerActionCtx string

const timerActionCtxKey TimerActionCtx = TimerActionCtx("timerActionCtxKey")

type Service interface {
	action.Trigger
}

type Handler struct {
	Service
}

type Model struct {
}

type Type int

const (
	MINUTE Type = iota + 1
	HOUR
	DAY
)

type ActionBody struct {
	Type     Type      `json:"type"`
	Name     string    `json:"name"`
	DateTime time.Time `json:"datetime"`
}
