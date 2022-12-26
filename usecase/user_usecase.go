// Package usecase describe business logic
package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/MochamadAkbar/go-restful/api"
	"github.com/MochamadAkbar/go-restful/common"
	"github.com/MochamadAkbar/go-restful/entity"
	"github.com/MochamadAkbar/go-restful/repository"
	"github.com/golang-jwt/jwt/v4"
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

func InitUseCase(repository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{
		Repository: repository,
	}
}

func (usecase *UserUsecase) Register(ctx context.Context, user *entity.User) (api.UserResponse, error) {
	var expiresIn = time.Now().Add(time.Duration(1) * time.Minute).Unix()
	
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

	token, err := common.JwtClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":   uuid,
			"expiresIn": expiresIn,
		},
		[]byte("487f17ff01c0753d6ef97593cbdeb476de7b5420c3006a4ac2c401b09187df7a"))
	if err != nil {
		panic(err)
	}

	resp.UserID = uuid
	resp.Token = token
	resp.ExpiresIn = expiresIn

	return resp, nil
}

func (usecase *UserUsecase) Login(ctx context.Context, user *entity.User) (api.UserResponse, error) {
	var expiresIn = time.Now().Add(time.Duration(1) * time.Minute).Unix()

	var resp api.UserResponse

	uuid, ok := usecase.Repository.Login(ctx, user)

	if !ok {
		return resp, errors.New("test")
	}

	token, err := common.JwtClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":    uuid,
			"expires_in": expiresIn,
		},
		[]byte("487f17ff01c0753d6ef97593cbdeb476de7b5420c3006a4ac2c401b09187df7a"))
	if err != nil {
		panic(err)
	}

	_, err = common.JwtValidate(common.HMAC, token, []byte("487f17ff01c0753d6ef97593cbdeb476de7b5420c3006a4ac2c401b09187df7a"))
	if err != nil {
		panic(err)
	}

	resp.UserID = uuid
	resp.Token = token
	resp.ExpiresIn = expiresIn

	return resp, nil
}
