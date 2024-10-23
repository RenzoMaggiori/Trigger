package reaction

import (
	"context"
	"net/http"

	"trigger.com/trigger/internal/action/workspace"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/jwt"
)

func (h *Handler) SendEmail(w http.ResponseWriter, r *http.Request) {
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

	err = h.Service.MutlipleReactions("send_email",
		context.WithValue(context.TODO(), AccessTokenCtxKey, token), actionNode)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}
