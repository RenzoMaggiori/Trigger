package session

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	errUserAlreadyExists error = errors.New("user already exists with provider")
	errSessionNotFound   error = errors.New("session not found")
	errBadSessionID      error = errors.New("bad session id")
	errBadUserID         error = errors.New("bad user id")
	errInsertSession     error = errors.New("could not insert session")

	errCodes map[error]customerror.CustomError = map[error]customerror.CustomError{
		errUserAlreadyExists: {
			Message: errUserAlreadyExists.Error(),
			Code:    http.StatusBadRequest,
		},
		errSessionNotFound: {
			Message: errSessionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errBadSessionID: {
			Message: errBadSessionID.Error(),
			Code:    http.StatusBadRequest,
		},
		errBadUserID: {
			Message: errBadSessionID.Error(),
			Code:    http.StatusBadRequest,
		},
		errInsertSession: {
			Message: errInsertSession.Error(),
			Code:    http.StatusInternalServerError,
		},
	}
)
