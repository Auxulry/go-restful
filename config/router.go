// Package config describe all configurations
package config

import (
	"github.com/MochamadAkbar/go-restful/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.CORS)
	router.Use(middleware.Recovery)

	return router
}
