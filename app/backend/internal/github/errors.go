package github

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	ErrSessionNotFound      error = errors.New("session not found")
	ErrUserNotFound         error = errors.New("user not found")
	ErrAccessTokenNotFound  error = errors.New("access token not found")
	ErrInvalidReactionInput error = errors.New("invalid reaction input")
	ErrInvalidReactionOuput error = errors.New("invalid reaction output")
	ErrInvalidGithubStatus  error = errors.New("invalid github status code received")

	ErrCodes map[error]customerror.CustomError = map[error]customerror.CustomError{
		ErrSessionNotFound: {
			Message: ErrSessionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrUserNotFound: {
			Message: ErrUserNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrAccessTokenNotFound: {
			Message: ErrAccessTokenNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		ErrInvalidReactionInput: {
			Message: ErrInvalidReactionOuput.Error(),
			Code:    http.StatusBadRequest,
		},
		ErrInvalidReactionOuput: {
			Message: ErrInvalidReactionOuput.Error(),
			Code:    http.StatusBadRequest,
		},
		ErrInvalidGithubStatus: {
			Message: ErrInvalidGithubStatus.Error(),
			Code:    http.StatusBadRequest,
		},
	}
)
