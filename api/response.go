// Package api describe all api request and response
package api

type Response struct {
	Code   int
	Status string
	Data   interface{}
}
