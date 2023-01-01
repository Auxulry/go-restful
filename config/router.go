// Package config describe all configurations
package config

import (
	"github.com/MochamadAkbar/go-restful/exception"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(exception.ErrorHandler)

	return router
}
