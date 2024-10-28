package trigger

import (
	"context"
	"log"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/jwt"
)

func (h *Handler) Watch(w http.ResponseWriter, r *http.Request) {

	token, err := jwt.FromRequest(r.Header.Get("Authorization"))
	if err != nil {
		customerror.Send(w, errors.ErrAuthorizationHeaderNotFound, errors.ErrCodes)
		return
	}

	actionNode, err := decode.Json[workspace.ActionNodeModel](r.Body)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	err = h.Service.Watch(context.WithValue(context.TODO(), AccessTokenCtxKey, token), actionNode)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}

func (h *Handler) Webhook(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming JSON body into the WebhookVerificationRequest struct
	webhookVerification, err := decode.Json[WebhookVerificationRequest](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	log.Printf("Webhook triggered, received body=%+v\n", webhookVerification)

	// Call the service's Webhook method
	err = h.Service.Webhook(context.TODO())
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	// Return the challenge string from the webhook verification request as a JSON response
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err = w.Write([]byte(webhookVerification.Challenge)); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) Stop(w http.ResponseWriter, r *http.Request) {

	token, err := jwt.FromRequest(r.Header.Get("Authorization"))
	if err != nil {
		customerror.Send(w, errors.ErrAuthorizationHeaderNotFound, errors.ErrCodes)
		return
	}

	err = h.Service.Stop(context.WithValue(context.TODO(), AccessTokenCtxKey, token))

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}
}
