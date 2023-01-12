// Package errors is utitlities for catch errors
package errors

import (
	"errors"
	"net/http"
)

type ErrHTTP struct {
	code    int
	message string
}

func (e *ErrHTTP) Code() int {
	return e.code
}

func (e *ErrHTTP) Error() string {
	return e.message
}

func NewErrHTTP(code int, message string) *ErrHTTP {
	return &ErrHTTP{
		code:    code,
		message: message,
	}
}

func GetHTTPCode(err error) int {
	var errHTTP *ErrHTTP
	if errors.As(err, &errHTTP) {
		return errHTTP.Code()
	} else {
		return http.StatusInternalServerError
	}
}
