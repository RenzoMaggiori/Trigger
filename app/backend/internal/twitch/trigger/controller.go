package trigger

import (
	"context"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
)

func (h *Handler) Watch(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Webhook(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	eventType := r.Header.Get("Twitch-Eventsub-Message-Type")
	webhookVerification, err := decode.Json[WebhookVerificationRequest](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	if eventType == VerificationMessageType {
		// Return the challenge string from the webhook verification request as a JSON response
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		if _, err = w.Write([]byte(webhookVerification.Challenge)); err != nil {
			customerror.Send(w, err, errors.ErrCodes)
			return
		}
		return
	}
	// Call the service's Webhook method
	err = h.Service.Webhook(context.WithValue(
		context.WithValue(
			r.Context(),
			WebhookUserIdCtxKey,
			userId,
		),
		WebhookVerificationCtxKey,
		webhookVerification),
	)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

}

func (h *Handler) Stop(w http.ResponseWriter, r *http.Request) {
	if err := h.Service.Stop(r.Context()); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}
