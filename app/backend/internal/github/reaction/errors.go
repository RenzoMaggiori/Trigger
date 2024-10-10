package reaction

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	errSessionNotFound      error = errors.New("session not found")
	errUserNotFound         error = errors.New("user not found")
	errAccessTokenNotFound  error = errors.New("access token not found")
	errInvalidReactionInput error = errors.New("invalid reaction input")

	errCodes map[error]customerror.CustomError = map[error]customerror.CustomError{
		errSessionNotFound: customerror.CustomError{
			Message: errSessionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errUserNotFound: customerror.CustomError{
			Message: errUserNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errAccessTokenNotFound: customerror.CustomError{
			Message: errAccessTokenNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errInvalidReactionInput: customerror.CustomError{
			Message: errInvalidReactionInput.Error(),
			Code:    http.StatusBadRequest,
		},
	}
)
