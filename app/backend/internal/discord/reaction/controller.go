package reaction

import (
	// "context"
	"log"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/middleware"
)

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	accessToken, ok := r.Context().Value(middleware.TokenCtxKey).(string)
	if !ok {
		customerror.Send(w, errors.ErrAccessTokenCtx, errors.ErrCodes)
		return
	}
	log.Println("accessToken: ", accessToken)

	actionNode, err := decode.Json[workspace.ActionNodeModel](r.Body)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	err = h.Service.MutlipleReactions("send_channel_message", r.Context(), actionNode)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}
