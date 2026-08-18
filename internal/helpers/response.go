package helpers

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, data any) error {
	return json.NewEncoder(w).Encode(data)
}

func ResponseString(w http.ResponseWriter, content string) error {
	_, err := w.Write([]byte(content))
	return err
}
