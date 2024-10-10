package trigger

import (
	"trigger.com/trigger/pkg/action"
)

type Service interface {
	action.Trigger
}

type Handler struct {
	Service
}

type Model struct {
}
