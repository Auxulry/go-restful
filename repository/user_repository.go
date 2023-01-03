// Package repository describe all action to db
package repository

import (
	"context"

	"github.com/MochamadAkbar/go-restful/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IUserRepository interface {
	Register(ctx context.Context, user *entity.User) bool
	Login(ctx context.Context, user *entity.User) (string, bool)
}

type UserRepository struct {
	Conn *pgxpool.Pool
}

func NewUserRepository(conn *pgxpool.Pool) IUserRepository {
	return &UserRepository{Conn: conn}
}

func (repository *UserRepository) Register(ctx context.Context, user *entity.User) bool {
	statement := `INSERT INTO "users" ("name", "email", "password") VALUES($1, $2, $3);`
	_, err := repository.Conn.Exec(ctx, statement, user.Name, user.Email, user.Password)
	if err != nil {
		panic(err.Error())
	}

	return true
}

func (repository *UserRepository) Login(ctx context.Context, user *entity.User) (string, bool) {
	statement := `SELECT "id" FROM "users" WHERE "email" = $1;`
	var ID string
	err := repository.Conn.QueryRow(ctx, statement, user.Email).
		Scan(&ID)
	if err != nil {
		panic(err.Error())
	}

	return ID, true
}
