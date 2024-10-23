package workspace

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	customerror "trigger.com/trigger/pkg/custom-error"
	"trigger.com/trigger/pkg/decode"
	"trigger.com/trigger/pkg/encode"
	"trigger.com/trigger/pkg/middleware"
)

func (h *Handler) GetWorkspace(w http.ResponseWriter, r *http.Request) {
	workspaces, err := h.Service.Get(context.TODO())

	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, workspaces); err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) GetWorkspaceById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errBadWorkspaceId, err)
		customerror.Send(w, error, errCodes)
		return
	}

	workspace, err := h.Service.GetById(context.TODO(), id)
	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, workspace); err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) GetWorkspacesByUserId(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("user_id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errBadUserId, err)
		customerror.Send(w, error, errCodes)
		return
	}

	workspaces, err := h.Service.GetByUserId(context.TODO(), id)
	if err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, workspaces); err != nil {
		log.Print(err)
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) AddWorkspace(w http.ResponseWriter, r *http.Request) {
	token, ok := r.Context().Value(middleware.TokenCtxKey).(string)
	log.Println("handler: ", token)
	if !ok {
		customerror.Send(w, errors.New("could not find token"), errCodes)
		return
	}

	add, err := decode.Json[AddWorkspaceModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	newWorkspace, err := h.Service.Add(context.WithValue(context.TODO(), AccessTokenCtxKey, token), &add)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, newWorkspace); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) ActionCompletedWorkspace(w http.ResponseWriter, r *http.Request) {

	update, err := decode.Json[ActionCompletedModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	accessToken := r.Header.Get("Authorization")

	updatedWorkspace, err := h.Service.ActionCompleted(
		context.WithValue(context.TODO(), AccessTokenCtxKey, accessToken),
		update,
	)

	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, updatedWorkspace); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}

func (h *Handler) UpdateById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

	if err != nil {
		error := fmt.Errorf("%w: %v", errBadWorkspaceId, err)
		customerror.Send(w, error, errCodes)
		return
	}

	update, err := decode.Json[UpdateWorkspaceModel](r.Body)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}

	updatedWorkspace, err := h.Service.UpdateById(context.TODO(), id, &update)
	if err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
	if err = encode.Json(w, updatedWorkspace); err != nil {
		customerror.Send(w, err, errCodes)
		return
	}
}

// func (h *Handler) DeleteById(w http.ResponseWriter, r *http.Request) {
// 	id, err := primitive.ObjectIDFromHex(r.PathValue("id"))

// 	if err != nil {
// 		error := fmt.Errorf("%w: %v", errBadActionId, err)
// 		customerror.Send(w, error, errCodes)
// 		return
// 	}
// 	if err := h.Service.DeleteById(context.TODO(), id); err != nil {
// 		customerror.Send(w, err, errCodes)
// 		return
// 	}
// }
