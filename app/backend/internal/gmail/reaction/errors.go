package reaction

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	errGmailSendEmail  error = errors.New("error while sending email through gmail")
	errSessionNotFound error = errors.New("session not found")
	errUserNotFound    error = errors.New("user not found")

	errCodes map[error]customerror.CustomError = map[error]customerror.CustomError{
		errGmailSendEmail: {
			Message: errGmailSendEmail.Error(),
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
	}
)
