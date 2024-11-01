package trigger

import (
	"context"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
)

func (h *Handler) WatchTimer(w http.ResponseWriter, r *http.Request) {
	actionNode, err := decode.Json[workspace.ActionNodeModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	err = h.Service.Watch(r.Context(), actionNode)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}

func (h *Handler) WebhookTimer(w http.ResponseWriter, r *http.Request) {
	event, err := decode.Json[ActionBody](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	err = h.Service.Webhook(
		context.WithValue(
			r.Context(),
			timerActionCtxKey,
			event,
		),
	)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) StopTimer(w http.ResponseWriter, r *http.Request) {
	err := h.Service.Stop(r.Context())
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}
