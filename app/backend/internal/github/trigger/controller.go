package trigger

import (
	"context"
	"net/http"

	githubClient "github.com/google/go-github/v66/github"
	"trigger.com/trigger/internal/action/workspace"
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

func (h *Handler) WebhookGithub(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	commit, err := decode.Json[githubClient.PushEvent](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	if err := h.Service.Webhook(
		context.WithValue(
			context.WithValue(
				r.Context(),
				GithubCommitCtxKey,
				commit,
			),
			userIdCtxKey,
			userId,
		),
	); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) StopGithub(w http.ResponseWriter, r *http.Request) {
	if err := h.Service.Stop(r.Context()); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}
