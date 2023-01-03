//go:build wireinject
// +build wireinject

// Package injector describe all dependency injectors
package injector

import (
	"net/http"

	"github.com/MochamadAkbar/go-restful/config"
	"github.com/MochamadAkbar/go-restful/handler"
	"github.com/MochamadAkbar/go-restful/repository"
	"github.com/MochamadAkbar/go-restful/usecase"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeServer(conn *pgxpool.Pool) *http.Server {
	wire.Build(
		config.NewRouter,
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		handler.NewUserHandler,
		config.NewServer,
	)

	return nil
}
