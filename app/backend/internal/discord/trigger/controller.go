package trigger

import (
	"context"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/errors"

	"trigger.com/trigger/pkg/decode"
)

func (h *Handler) WatchDiscord(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) WebhookDiscord(w http.ResponseWriter, r *http.Request) {
	event, err := decode.Json[ActionBody](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}

	err = h.Service.Webhook(context.WithValue(r.Context(), DiscordEventCtxKey, event))
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) StopDiscord(w http.ResponseWriter, r *http.Request) {
	err := h.Service.Stop(r.Context())
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}
