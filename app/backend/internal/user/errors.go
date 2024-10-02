package user

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	errUserAlreadyExists error = errors.New("user already exists")
	errUserNotFound      error = errors.New("user not found")

	errCodes map[error]customerror.CustomError = map[error]customerror.CustomError{
		errUserAlreadyExists: {
			Message: errUserAlreadyExists.Error(),
			Code:    http.StatusBadRequest,
		},
		errUserNotFound: {
			Message: errUserNotFound.Error(),
			Code:    http.StatusNotFound,
		},
	}
)
