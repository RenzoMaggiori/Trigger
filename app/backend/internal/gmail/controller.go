package gmail

import (
	"context"
	"log"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
)

func (h *Handler) WatchGmail(w http.ResponseWriter, r *http.Request) {
	log.Printf("Watching gmail")
	accessToken := r.Header.Get("Authorization")
	actionNode, err := decode.Json[workspace.ActionNodeModel](r.Body)

	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	err = h.Service.Watch(context.WithValue(context.TODO(), AccessTokenCtxKey, accessToken), actionNode)

	if err != nil {
		customerror.Send(w, err, errCodes)
	}
}

func (h *Handler) WebhookGmail(w http.ResponseWriter, r *http.Request) {

	body, err := decode.Json[Event](r.Body)

	if err != nil {
		customerror.Send(w, err, errCodes)
	}

	log.Printf("Webhook triggered, received body=%+v\n", body)

	err = h.Service.Webhook(context.TODO())

	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}
