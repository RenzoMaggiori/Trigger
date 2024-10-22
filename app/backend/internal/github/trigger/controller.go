package trigger

import (
	"context"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/github"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/errors"

	"trigger.com/trigger/pkg/decode"
)

func (h *Handler) WatchGithub(w http.ResponseWriter, r *http.Request) {
	actionNode, err := decode.Json[workspace.ActionNodeModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	if err := h.Service.Watch(r.Context(), actionNode); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) StopGithub(w http.ResponseWriter, r *http.Request) {
	body, err := decode.Json[StopModel](r.Body)
	if err != nil {
		customerror.Send(w, errors.ErrGithubStopModelNotFound, errors.ErrCodes)
		return
	}

	if err := h.Service.Stop(
		context.WithValue(
			r.Context(),
			github.StopCtxKey,
			body,
		),
	); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) WebhookGithub(w http.ResponseWriter, r *http.Request) {
	if err := h.Service.Webhook(r.Context()); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}
