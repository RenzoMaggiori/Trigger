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
	errUserNotCreated   error = errors.New("could not create user")
	errUserNotUpdated   error = errors.New("could not uptade user")
	errUserNotDeleted   error = errors.New("could not delete user")

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
		errUserNotCreated: {
			Message: errUserNotCreated.Error(),
			Code: http.StatusBadRequest,
		},
		errUserNotUpdated: {
			Message: errUserNotUpdated.Error(),
			Code: http.StatusBadRequest,
		},
		errUserNotDeleted: {
			Message: errUserNotDeleted.Error(),
			Code: http.StatusBadRequest,
		},
	}
)
