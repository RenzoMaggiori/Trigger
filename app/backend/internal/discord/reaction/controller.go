package reaction

import (
	// "context"
	"log"
	"net/http"

	// "trigger.com/trigger/internal/action/workspace"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
)

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Sending message")

	// accessToken := r.Header.Get("Authorization")
	// actionNode, err := decode.Json[workspace.ActionNodeModel](r.Body)
	// if err != nil {
	// 	customerror.Send(w, err, errors.ErrCodes)
	// 	return
	// }

	newMsg, err := decode.Json[NewMessage](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	err = h.Service.SendMessage(newMsg)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}
