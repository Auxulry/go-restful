// Package usecase describe business logic
package usecase

import (
	"context"
	"errors"

	"github.com/MochamadAkbar/go-restful/api"
	"github.com/MochamadAkbar/go-restful/entity"
	"github.com/MochamadAkbar/go-restful/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	Salt = 4
)

type IUserUsecase interface {
	Register(ctx context.Context, user *entity.User) (api.UserResponse, error)
	Login(ctx context.Context, user *entity.User) (api.UserResponse, error)
}

type UserUsecase struct {
	Repository repository.IUserRepository
}

func (usecase *UserUsecase) Register(ctx context.Context, user *entity.User) (api.UserResponse, error) {
	var resp api.UserResponse

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), Salt)
	if err != nil {
		return resp, err
	}

	user.Password = string(hash)

	ok := usecase.Repository.Register(ctx, user)

	if !ok {
		return resp, errors.New("test")
	}

	uuid, ok := usecase.Repository.Login(ctx, user)

	if !ok {
		return resp, errors.New("test")
	}

	resp.UUID = uuid
	resp.Token = "Token"

	return resp, nil
}

func (usecase *UserUsecase) Login(ctx context.Context, user *entity.User) (api.UserResponse, error) {
	var resp api.UserResponse
	uuid, ok := usecase.Repository.Login(ctx, user)

	if !ok {
		return resp, errors.New("test")
	}

	resp.UUID = uuid
	resp.Token = "Token"

	return resp, nil
}
