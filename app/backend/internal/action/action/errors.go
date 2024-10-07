package action

import (
	"errors"
	"net/http"

	customerror "trigger.com/trigger/pkg/custom-error"
)

var (
	errActionNotFound error                             = errors.New("action not found")
	errBadActionId    error                             = errors.New("bad action id")
	errBadActionType  error                             = errors.New("bad action type")
	errCodes          map[error]customerror.CustomError = map[error]customerror.CustomError{

		errActionNotFound: {
			Message: errActionNotFound.Error(),
			Code:    http.StatusNotFound,
		},
		errBadActionId: {
			Message: errBadActionId.Error(),
			Code:    http.StatusBadRequest,
		},
		errBadActionType: {
			Message: errBadActionType.Error(),
			Code:    http.StatusBadRequest,
		},
	}
)
