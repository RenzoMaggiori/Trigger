package decode

import (
	"encoding/json"
	"io"
)

func Json[T any](r io.Reader) (T, error) {
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, err
	}
	return v, nil
}
