// Package repository describe all action to db
package repository

import (
	"context"
	"github.com/MochamadAkbar/go-restful/api"
	"github.com/MochamadAkbar/go-restful/entity"
)

type UserRepository interface {
	Register(ctx context.Context, user *entity.User) (*entity.User, error)
	Login(ctx context.Context, user *api.UserRequest) (*entity.User, error)
}
