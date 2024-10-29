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

func (h *Handler) WatchGmail(w http.ResponseWriter, r *http.Request) {

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

func (h *Handler) WebhookGmail(w http.ResponseWriter, r *http.Request) {

	event, err := decode.Json[Event](r.Body)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
	}

	// log.Printf("Webhook triggered, received body=%+v\n", event)

	err = h.Service.Webhook(context.WithValue(context.TODO(), GmailEventCtxKey, event))

	if err != nil {
		log.Println(err)
		return
	}
}

func (h *Handler) StopGmail(w http.ResponseWriter, r *http.Request) {

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
