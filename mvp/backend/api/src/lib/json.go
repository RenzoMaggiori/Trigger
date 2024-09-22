package lib

import (
	"encoding/json"
	"io"
)

func JsonDecode(r io.Reader) (any, error) {
	var v any
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}
