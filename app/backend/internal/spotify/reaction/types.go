package reaction

import (
	"trigger.com/trigger/pkg/action"
)

type Service interface {
	action.MultipleReactions
}

type Handler struct {
	Service Service
}

type Model struct {
}
