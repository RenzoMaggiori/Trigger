package user

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	errUserAlreadyExists error = errors.New("user already exists")
	errUserNotFound      error = errors.New("user not found")
	errBadUserID         error = errors.New("bad user id")

	errCodes map[error]customerror.CustomError = map[error]customerror.CustomError{
		errUserAlreadyExists: {
			Message: errUserAlreadyExists.Error(),
			Code:    http.StatusConflict,
		},
		errUserNotFound: {
			Message: errUserNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errBadUserID: {
			Message: errBadUserID.Error(),
			Code:    http.StatusBadRequest,
		},
	}
)
