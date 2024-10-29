package workspace

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"trigger.com/trigger/internal/session"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/encode"
	"trigger.com/trigger/pkg/errors"
	"trigger.com/trigger/pkg/jwt"
	"trigger.com/trigger/pkg/middleware"
)

func (h *Handler) GetWorkspace(w http.ResponseWriter, r *http.Request) {
	workspaces, err := h.Service.Get(context.TODO())

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

func (h *Handler) GetMyWorkspaces(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetWorkspaceById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errors.ErrBadWorkspaceId, err)
		customerror.Send(w, error, errors.ErrCodes)
		return
	}

	workspace, err := h.Service.GetById(context.TODO(), id)
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

func (h *Handler) GetWorkspacesByUserId(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("user_id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errors.ErrBadUserId, err)
		customerror.Send(w, error, errors.ErrCodes)
		return
	}

	workspaces, err := h.Service.GetByUserId(context.TODO(), id)
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

func (h *Handler) GetWorkspacesByActionId(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("action_id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errors.ErrBadActionId, err)
		customerror.Send(w, error, errors.ErrCodes)
		return
	}

	workspaces, err := h.Service.GetByActionId(context.TODO(), id)
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

func (h *Handler) AddWorkspace(w http.ResponseWriter, r *http.Request) {
	token, ok := r.Context().Value(middleware.TokenCtxKey).(string)

	if !ok {
		customerror.Send(w, errors.ErrAccessTokenCtx, errors.ErrCodes)
		return
	}

	add, err := decode.Json[AddWorkspaceModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	newWorkspace, err := h.Service.Add(context.WithValue(context.TODO(), AccessTokenCtxKey, token), &add)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	if err = encode.Json(w, newWorkspace); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

func (h *Handler) ActionCompletedWorkspace(w http.ResponseWriter, r *http.Request) {
	update, err := decode.Json[ActionCompletedModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	token, err := jwt.FromRequest(r.Header.Get("Authorization"))

	if err != nil {
		customerror.Send(w, errors.ErrAuthorizationHeaderNotFound, errors.ErrCodes)
		return
	}

	updatedWorkspace, err := h.Service.ActionCompleted(
		context.WithValue(context.TODO(), AccessTokenCtxKey, token),
		update,
	)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	if err = encode.Json(w, updatedWorkspace); err != nil {
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

	updatedWorkspace, err := h.Service.UpdateById(context.TODO(), id, &update)
	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
	if err = encode.Json(w, updatedWorkspace); err != nil {
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

	updatedWorkspace, err := h.Service.WatchCompleted(
		context.TODO(),
		update,
	)

	if err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}

	if err = encode.Json(w, updatedWorkspace); err != nil {
		customerror.Send(w, err, errors.ErrCodes)
		return
	}
}

// func (h *Handler) DeleteById(w http.ResponseWriter, r *http.Request) {
// 	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

// 	if err != nil {
// 		error := fmt.Errorf("%w: %v", errBadActionId, err)
// 		customerror.Send(w, error, errors.ErrCodes)
// 		return
// 	}
// 	if err := h.Service.DeleteById(context.TODO(), id); err != nil {
// 		customerror.Send(w, err, errors.ErrCodes)
// 		return
// 	}
// }
