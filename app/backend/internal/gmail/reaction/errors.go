package reaction

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	errGmailSendEmail error = errors.New("error while sending email through gmail")

	errCodes map[error]customerror.CustomError = map[error]customerror.CustomError{
		errGmailSendEmail: {
			Message: errGmailSendEmail.Error(),
			Code:    http.StatusNotFound,
		},
	}
)
