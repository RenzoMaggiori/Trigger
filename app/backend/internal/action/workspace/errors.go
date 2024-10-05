package workspace

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	errWorkspaceNotFound  error                             = errors.New("action not found")
	errBadWorkspaceId     error                             = errors.New("bad workspace id")
	errBadUserId          error                             = errors.New("bad user id")
	errSessionNotFound    error                             = errors.New("session not found")
	errSessionTypeNone    error                             = errors.New("could not decypher session type")
	errCreatingWorkspace  error                             = errors.New("error while creating workspace")
	errFetchingActions    error                             = errors.New("could not fetch actions")
	errActionTypeNone     error                             = errors.New("could not decypher action type")
	errActionWatch        error                             = errors.New("action watch failed")
	errActionNodeTypeNone error                             = errors.New("could not decypher action node type")
	errCodes              map[error]customerror.CustomError = map[error]customerror.CustomError{
		errWorkspaceNotFound: {
			Message: errWorkspaceNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errBadWorkspaceId: {
			Message: errBadWorkspaceId.Error(),
			Code:    http.StatusBadRequest,
		},
		errBadUserId: {
			Message: errBadUserId.Error(),
			Code:    http.StatusBadRequest,
		},
		errSessionNotFound: {
			Message: errSessionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errSessionTypeNone: {
			Message: errSessionTypeNone.Error(),
			Code:    http.StatusNotFound,
		},
		errCreatingWorkspace: {
			Message: errCreatingWorkspace.Error(),
			Code:    http.StatusInternalServerError,
		},
		errFetchingActions: {
			Message: errFetchingActions.Error(),
			Code:    http.StatusInternalServerError,
		},
		errActionTypeNone: {
			Message: errActionTypeNone.Error(),
			Code:    http.StatusNotFound,
		},
		errActionWatch: {
			Message: errActionWatch.Error(),
			Code:    http.StatusInternalServerError,
		},
		errActionNodeTypeNone: {
			Message: errActionNodeTypeNone.Error(),
			Code:    http.StatusInternalServerError,
		},
	}
)
