package customerror

import (
	"errors"
	"log"
	"net/http"
)

type CustomError struct {
	Message string
	Code    int
}

func get(err error, customs map[error]CustomError) *CustomError {
	for e, custom := range customs {
		if errors.Is(err, e) {
			return &custom
		}
	}
	return &CustomError{
		Message: "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func Send(w http.ResponseWriter, err error, customs map[error]CustomError) {
	log.Println(err)
	cerr := get(err, customs)
	http.Error(w, cerr.Message, cerr.Code)
}
