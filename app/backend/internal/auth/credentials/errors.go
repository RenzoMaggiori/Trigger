package credentials

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	errCreateuser                  error = errors.New("unable to create user")
	errCredentialsNotFound         error = errors.New("could not get credentials from context")
	errAuthorizationHeaderNotFound error = errors.New("could not get authorization header")
	errAuthorizationTypeNone       error = errors.New("could not decypher auth type")
	errTokenNotFound               error = errors.New("could not find token in authorization header")
	errAuthTypeUndefined           error = errors.New("auth type is undefined")
	errUserNotRetrieved            error = errors.New("could not retrieve user")
	errSessionNotRetrieved         error = errors.New("could not retrieve session")
	errCreateToken                 error = errors.New("unable to create token")
	errSessionNotFound             error = errors.New("session not found")
	errCreateSession               error = errors.New("unable to create session")

	errCodes map[error]customerror.CustomError = map[error]customerror.CustomError{
		errCredentialsNotFound: {
			Message: errCredentialsNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errAuthorizationHeaderNotFound: {
			Message: errAuthorizationHeaderNotFound.Error(),
			Code:    http.StatusUnprocessableEntity,
		},
		errAuthorizationTypeNone: {
			Message: errAuthorizationTypeNone.Error(),
			Code:    http.StatusNotFound,
		},
		errTokenNotFound: {
			Message: errTokenNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errAuthTypeUndefined: {
			Message: errAuthTypeUndefined.Error(),
			Code:    http.StatusUnprocessableEntity,
		},
		errCreateuser: {
			Message: errCreateuser.Error(),
			Code:    http.StatusBadRequest,
		},
		errSessionNotRetrieved: {
			Message: errSessionNotRetrieved.Error(),
			Code:    http.StatusNotFound,
		},
		errUserNotRetrieved: {
			Message: errUserNotRetrieved.Error(),
			Code:    http.StatusNotFound,
		},
		errCreateToken: {
			Message: errCreateToken.Error(),
			Code:    http.StatusBadRequest,
		},
		errSessionNotFound: {
			Message: errSessionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errCreateSession: {
			Message: errCreateSession.Error(),
			Code:    http.StatusBadRequest,
		},
	}
)
