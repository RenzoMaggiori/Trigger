package trigger

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	errGmailWatch error = errors.New("gmail watch failed")

	errCodes map[error]customerror.CustomError = map[error]customerror.CustomError{
		errGmailWatch: {
			Message: errGmailWatch.Error(),
			Code:    http.StatusNotFound,
		},
	}
)
