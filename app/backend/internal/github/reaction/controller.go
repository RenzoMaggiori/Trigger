package reaction

import (
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	"trigger.com/trigger/internal/github"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
)

func (h *Handler) CreateIssue(w http.ResponseWriter, r *http.Request) {
	actionNode, err := decode.Json[workspace.ActionNodeModel](r.Body)
	if err != nil {
		customerror.Send(w, err, github.ErrCodes)
		return
	}

	if err := h.Service.Reaction(r.Context(), actionNode); err != nil {
		customerror.Send(w, err, github.ErrCodes)
		return
	}
}
