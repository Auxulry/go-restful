//go:build wireinject
// +build wireinject

// Package injector describe all dependency injectors
package injector

import (
	"github.com/MochamadAkbar/go-restful/handler"
	"github.com/MochamadAkbar/go-restful/repository"
	"github.com/MochamadAkbar/go-restful/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeUserService(conn *pgxpool.Pool, router chi.Router) error {
	wire.Build(
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		handler.NewUserHandler,
	)

	return nil
}
