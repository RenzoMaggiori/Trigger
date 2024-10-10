package trigger

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	errGmailWatch      error                             = errors.New("error while watching gmail")
	errGmailStop       error                             = errors.New("error while stropping gmail")
	errSessionNotFound error                             = errors.New("session not found")
	errUserNotFound    error                             = errors.New("user not found")
	errEventCtx        error                             = errors.New("could not retrieve event")
	errActionNotFound  error                             = errors.New("could not retrieve action")
	errCodes           map[error]customerror.CustomError = map[error]customerror.CustomError{
		errGmailWatch: {
			Message: errGmailWatch.Error(),
			Code:    http.StatusNotFound,
		},
		errGmailStop: {
			Message: errGmailStop.Error(),
			Code:    http.StatusNotFound,
		},
		errSessionNotFound: {
			Message: errSessionNotFound.Error(),
			Code:    http.StatusInternalServerError,
		},
		errUserNotFound: {
			Message: errUserNotFound.Error(),
			Code:    http.StatusInternalServerError,
		},
		errEventCtx: {
			Message: errEventCtx.Error(),
			Code:    http.StatusInternalServerError,
		},
		errActionNotFound: {
			Message: errActionNotFound.Error(),
			Code:    http.StatusInternalServerError,
		},
	}
)
