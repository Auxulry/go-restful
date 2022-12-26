package common

import (
	"encoding/json"
	"net/http"
)

func SerializeRequest(r *http.Request, result interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func SerializeWriter(w http.ResponseWriter, code int, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		return err
	}
	return nil
}
