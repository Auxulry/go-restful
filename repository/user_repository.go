// Package repository describe all action to db
package repository

import (
	"context"

	"github.com/MochamadAkbar/go-restful/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Register(ctx context.Context, user *entity.User) (entity.User, error)
	Login(ctx context.Context, user *entity.User) (entity.User, error)
}

type UserRepositoryImpl struct {
	Conn *pgxpool.Pool
}

func NewUserRepository(conn *pgxpool.Pool) UserRepository {
	return &UserRepositoryImpl{Conn: conn}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, user *entity.User) (entity.User, error) {
	var row entity.User
	statement := `INSERT INTO "users" ("name", "email", "password") VALUES($1, $2, $3) RETURNING "id";`
	err := repository.Conn.QueryRow(ctx, statement, user.Name, user.Email, user.Password).Scan(&row.ID)
	if err != nil {
		return row, err
	}

	return row, nil
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, user *entity.User) (entity.User, error) {
	var row entity.User
	statement := `SELECT "id", "email", "password" FROM "users" WHERE "email" = $1;`
	err := repository.Conn.QueryRow(ctx, statement, user.Email).
		Scan(&row.ID, &row.Email, &row.Password)
	if err != nil {
		return row, err
	}

	return row, nil
}
