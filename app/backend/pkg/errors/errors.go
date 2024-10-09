package errors

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	ErrWorkspaceNotFound   error                             = errors.New("action not found")
	ErrBadWorkspaceId      error                             = errors.New("bad workspace id")
	ErrBadUserId           error                             = errors.New("bad user id")
	ErrSessionNotFound     error                             = errors.New("session not found")
	ErrSessionTypeNone     error                             = errors.New("could not decypher session type")
	ErrCreatingWorkspace   error                             = errors.New("error while creating workspace")
	ErrFetchingActions     error                             = errors.New("could not fetch actions")
	ErrActionTypeNone      error                             = errors.New("could not decypher action type")
	ErrAction              error                             = errors.New("action service failed")
	ErrActionNotFound      error                             = errors.New("action not found")
	ErrActionNodeTypeNone  error                             = errors.New("could not decypher action node type")
	ErrUserNotFound        error                             = errors.New("user not found")
	ErrUserTypeNone        error                             = errors.New("could not decypher user type")
	ErrSettingAction       error                             = errors.New("error while setting trigger or reaction")
	ErrCompletingAction    error                             = errors.New("error while completing action")
	ErrSessionNotRetrieved error                             = errors.New("error while retrieving session")
	ErrSessionNotCreated   error                             = errors.New("error while creating session")
	ErrAccessTokenCtxKey   error                             = errors.New("could not retrieve access token from context")
	ErrFailedToCreateEmail error                             = errors.New("failed to create raw email")
	ErrFailedToSendEmail   error                             = errors.New("failed to send email")
	ErrCodes               map[error]customerror.CustomError = map[error]customerror.CustomError{
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
		ErrUserTypeNone: {
			Message: ErrUserTypeNone.Error(),
			Code:    http.StatusNotFound,
		},
		ErrSettingAction: {
			Message: ErrSettingAction.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrCompletingAction: {
			Message: ErrCompletingAction.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrActionNotFound: {
			Message: ErrActionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrSessionNotRetrieved: {
			Message: ErrSessionNotRetrieved.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrSessionNotCreated: {
			Message: ErrSessionNotCreated.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrAccessTokenCtxKey: {
			Message: ErrAccessTokenCtxKey.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrFailedToCreateEmail: {
			Message: ErrFailedToCreateEmail.Error(),
			Code:    http.StatusInternalServerError,
		},
		ErrFailedToSendEmail: {
			Message: ErrFailedToSendEmail.Error(),
			Code:    http.StatusInternalServerError,
		},
	}
)
