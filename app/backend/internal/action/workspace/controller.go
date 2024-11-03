package workspace

import (
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/internal/session"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/encode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/middleware"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	workspaces, err := h.Service.Get(r.Context())

	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	if err = encode.Json(w, workspaces); err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) GetByAcessToken(w http.ResponseWriter, r *http.Request) {
	token, ok := r.Context().Value(middleware.TokenCtxKey).(string)
	if !ok {
		customerror.Send(w, errors.ErrAccessTokenCtx, errors.ErrCodes)
		return
	}

	s, _, err := session.GetSessionByAccessTokenRequest(token)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	workspaces, err := h.Service.GetByUserId(r.Context(), s.UserId)
	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	if err = encode.Json(w, workspaces); err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errors.ErrBadWorkspaceId, err)
		customerror.Send(w, error, errors.ErrCodes)
		return
	}

	workspace, err := h.Service.GetById(r.Context(), id)
	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	if err = encode.Json(w, workspace); err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) GetByUserId(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("user_id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errors.ErrBadUserId, err)
		customerror.Send(w, error, errors.ErrCodes)
		return
	}

	workspaces, err := h.Service.GetByUserId(r.Context(), id)
	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	if err = encode.Json(w, workspaces); err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) GetByActionId(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("action_id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errors.ErrBadActionId, err)
		customerror.Send(w, error, errors.ErrCodes)
		return
	}

	workspaces, err := h.Service.GetByActionId(r.Context(), id)
	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	if err = encode.Json(w, workspaces); err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	add, err := decode.Json[AddWorkspaceModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	newWorkspace, err := h.Service.Add(r.Context(), &add)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	if err = encode.Json(w, newWorkspace); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) UpdateById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errors.ErrBadWorkspaceId, err)
		customerror.Send(w, error, errors.ErrCodes)
		return
	}

	update, err := decode.Json[UpdateWorkspaceModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	updatedWorkspace, err := h.Service.UpdateById(r.Context(), id, &update)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	if err = encode.Json(w, updatedWorkspace); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) ActionCompleted(w http.ResponseWriter, r *http.Request) {
	update, err := decode.Json[ActionCompletedModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	err = h.Service.ActionCompleted(
		r.Context(),
		update,
	)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) WatchCompleted(w http.ResponseWriter, r *http.Request) {
	update, err := decode.Json[WatchCompletedModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	err = h.Service.WatchCompleted(
		r.Context(),
		update,
	)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

}

func (h *Handler) Start(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errors.ErrBadWorkspaceId, err)
		customerror.Send(w, error, errors.ErrCodes)
		return
	}

	workspace, err := h.Service.Start(r.Context(), id)
	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	if err = encode.Json(w, workspace); err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) Stop(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errors.ErrBadWorkspaceId, err)
		customerror.Send(w, error, errors.ErrCodes)
		return
	}

	workspace, err := h.Service.Stop(r.Context(), id)
	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	if err = encode.Json(w, workspace); err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errors.ErrBadWorkspaceId, err)
		customerror.Send(w, error, errors.ErrCodes)
		return
	}

	if err := h.Service.DeleteById(r.Context(), id); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) GetTemplates(w http.ResponseWriter, r *http.Request) {
	templates, err := h.Service.Templates(r.Context())
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	if err = encode.Json(w, templates); err != nil {
		log.Print(err)
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}
