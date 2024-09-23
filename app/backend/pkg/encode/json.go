package encode

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data any) error {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return nil
}
