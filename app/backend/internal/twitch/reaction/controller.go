package reaction

import (
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
)

func (h *Handler) SendChannelMessage(w http.ResponseWriter, r *http.Request) {
	actionNode, err := decode.Json[workspace.ActionNodeModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	err = h.Service.MutlipleReactions("send_chat_message", r.Context(), actionNode)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}
