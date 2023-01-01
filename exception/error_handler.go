// Package exception describe all exception utils
package exception

import (
	"net/http"

	"github.com/MochamadAkbar/go-restful/api"
	"github.com/MochamadAkbar/go-restful/common"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if errRecover := recover(); errRecover != nil {
				internalServerError(w, r, errRecover)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func internalServerError(w http.ResponseWriter, _ *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	errResponse := api.ErrResponse{
		Code:    http.StatusInternalServerError,
		Status:  http.StatusText(http.StatusInternalServerError),
		Message: err,
	}

	_ = common.SerializeWriter(w, http.StatusInternalServerError, errResponse)
}
