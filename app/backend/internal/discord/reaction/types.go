package reaction

import (
	"trigger.com/trigger/pkg/action"
)

type Service interface {
	action.Reaction
}

type Handler struct {
	Service
}

type Model struct {
}
