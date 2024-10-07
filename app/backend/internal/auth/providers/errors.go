package providers

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	errCredentialsNotFound         error = errors.New("could not get credentials from context")
	errAuthorizationHeaderNotFound error = errors.New("could not get authorization header")
	errAuthorizationTypeNone       error = errors.New("could not decypher auth type")
	errTokenNotFound               error = errors.New("could not find/verify token")
	errAuthTypeUndefined           error = errors.New("auth type is undefined")
	errUserNotFound                error = errors.New("could not find user")
	errUserTypeNone                error = errors.New("could not decypher user type")
	errSessionNotFound             error = errors.New("could not find session")
	errSessionTypeNone             error = errors.New("could not decypher session type")
	errProviderSessionNotFound     error = errors.New("could not find provider session")
	errSessionPatchFailed          error = errors.New("could not patch session")

	errCodes map[error]customerror.CustomError = map[error]customerror.CustomError{
		errCredentialsNotFound: {
			Message: errCredentialsNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errAuthorizationHeaderNotFound: {
			Message: errAuthorizationHeaderNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errAuthorizationTypeNone: {
			Message: errAuthorizationTypeNone.Error(),
			Code:    http.StatusInternalServerError,
		},
		errTokenNotFound: {
			Message: errTokenNotFound.Error(),
			Code:    http.StatusUnauthorized,
		},
		errAuthTypeUndefined: {
			Message: errAuthTypeUndefined.Error(),
			Code:    http.StatusNotFound,
		},
		errSessionNotFound: {
			Message: errSessionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errUserNotFound: {
			Message: errUserNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errUserTypeNone: {
			Message: errUserTypeNone.Error(),
			Code:    http.StatusNotFound,
		},
		errSessionTypeNone: {
			Message: errSessionTypeNone.Error(),
			Code:    http.StatusNotFound,
		},
		errProviderSessionNotFound: {
			Message: errProviderSessionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errSessionPatchFailed: {
			Message: errSessionPatchFailed.Error(),
			Code:    http.StatusInternalServerError,
		},
	}
)
