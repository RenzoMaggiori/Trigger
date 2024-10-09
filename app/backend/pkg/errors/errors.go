package errors

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	ErrWorkspaceNotFound  error = errors.New("action not found")
	ErrBadWorkspaceId     error = errors.New("bad workspace id")
	ErrBadUserId          error = errors.New("bad user id")
	ErrSessionNotFound    error = errors.New("session not found")
	ErrSessionTypeNone    error = errors.New("could not decypher session type")
	ErrCreatingWorkspace  error = errors.New("error while creating workspace")
	ErrFetchingActions    error = errors.New("could not fetch actions")
	ErrActionTypeNone     error = errors.New("could not decypher action type")
	ErrAction             error = errors.New("action service failed")
	ErrActionNodeTypeNone error = errors.New("could not decypher action node type")
	ErrUserNotFound       error = errors.New("user not found")

	ErrCodes map[error]customerror.CustomError = map[error]customerror.CustomError{
		ErrWorkspaceNotFound: {
			Message: ErrWorkspaceNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrBadWorkspaceId: {
			Message: ErrBadWorkspaceId.Error(),
			Code:    http.StatusBadRequest,
		},
		ErrBadUserId: {
			Message: ErrBadUserId.Error(),
			Code:    http.StatusBadRequest,
		},
		ErrSessionNotFound: {
			Message: ErrSessionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrSessionTypeNone: {
			Message: ErrSessionTypeNone.Error(),
			Code:    http.StatusNotFound,
		},
		ErrCreatingWorkspace: {
			Message: ErrCreatingWorkspace.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrFetchingActions: {
			Message: ErrFetchingActions.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrActionTypeNone: {
			Message: ErrActionTypeNone.Error(),
			Code:    http.StatusNotFound,
		},
		ErrActionNodeTypeNone: {
			Message: ErrActionNodeTypeNone.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrAction: {
			Message: ErrAction.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrUserNotFound: {
			Message: ErrUserNotFound.Error(),
			Code:    http.StatusNotFound,
		},
	}
)
