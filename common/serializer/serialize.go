// Package serializer is utilities to decode and encode json
package serializer

import (
	"encoding/json"
	"net/http"
)

func SerializeRequest(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	if err != nil {
		panic(err)
	}
}

func SerializeWriter(w http.ResponseWriter, code int, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		panic(err)
	}
}

func SerializeCloseRequest(r *http.Request) {
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
}
