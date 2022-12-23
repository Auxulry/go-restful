// Package repository describe all action to db
package repository

import (
	"context"

	"github.com/MochamadAkbar/go-restful/api"
	"github.com/MochamadAkbar/go-restful/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IUserRepository interface {
	Register(ctx context.Context, user *entity.User) bool
	Login(ctx context.Context, user *api.UserRequest) (string, bool)
}

type UserRepository struct {
	Conn *pgxpool.Conn
}

func (repository UserRepository) Register(ctx context.Context, user *entity.User) bool {
	statement := `INSERT INTO "users" ("name", "email", "password") VALUES($1, $2, $3);`
	_, err := repository.Conn.Exec(ctx, statement, user.Name, user.Email, user.Password)
	if err != nil {
		panic(err)
	}

	return true
}

func (repository UserRepository) Login(ctx context.Context, user *api.UserRequest) (string, bool) {
	statement := `SELECT "uuid" FROM "users" WHERE "email" = $1;`
	var UUID string
	err := repository.Conn.QueryRow(ctx, statement, user.Email).
		Scan(&UUID)
	if err != nil {
		panic(err)
	}

	return UUID, true
}
