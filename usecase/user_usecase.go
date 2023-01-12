// Package usecase describe business logic
package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/MochamadAkbar/go-restful/api"
	commonErr "github.com/MochamadAkbar/go-restful/common/errors"
	commonJwt "github.com/MochamadAkbar/go-restful/common/jwt"
	"github.com/MochamadAkbar/go-restful/entity"
	"github.com/MochamadAkbar/go-restful/exception"
	"github.com/MochamadAkbar/go-restful/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	Salt = 4
)

type UserUsecase interface {
	Register(ctx context.Context, user *entity.User) (api.UserResponse, error)
	Login(ctx context.Context, user *entity.User) (api.UserResponse, error)
}

type UserUsecaseImpl struct {
	Repository repository.UserRepository
}

func NewUserUsecase(repository repository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{
		Repository: repository,
	}
}

func (usecase *UserUsecaseImpl) Register(ctx context.Context, user *entity.User) (api.UserResponse, error) {
	var resp api.UserResponse
	expiresIn := time.Now().Add(time.Duration(1) * time.Minute).Unix()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), Salt)
	if err != nil {
		return resp, commonErr.NewErrHTTP(http.StatusInternalServerError, err.Error())
	}

	user.Password = string(hash)

	row, err := usecase.Repository.Register(ctx, user)
	if err != nil {
		return resp, exception.NewException(err.Error())
	}

	token, err := commonJwt.JwtClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":   row.ID,
			"expiresIn": expiresIn,
		},
		[]byte("487f17ff01c0753d6ef97593cbdeb476de7b5420c3006a4ac2c401b09187df7a"))
	if err != nil {
		return resp, commonErr.NewErrHTTP(http.StatusInternalServerError, err.Error())
	}

	resp.UserID = row.ID
	resp.Token = token
	resp.ExpiresIn = expiresIn

	return resp, nil
}

func (usecase *UserUsecaseImpl) Login(ctx context.Context, user *entity.User) (api.UserResponse, error) {
	var resp api.UserResponse
	expiresIn := time.Now().Add(time.Duration(1) * time.Minute).Unix()

	row, err := usecase.Repository.Login(ctx, user)
	if err != nil {
		return resp, commonErr.NewErrHTTP(http.StatusNotFound, "User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(row.Password), []byte(user.Password))
	if err != nil {
		return resp, commonErr.NewErrHTTP(http.StatusUnauthorized, "Invalid Credentials")
	}
	token, err := commonJwt.JwtClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":    row.ID,
			"expires_in": expiresIn,
		},
		[]byte("487f17ff01c0753d6ef97593cbdeb476de7b5420c3006a4ac2c401b09187df7a"))
	if err != nil {
		return resp, commonErr.NewErrHTTP(http.StatusInternalServerError, err.Error())
	}

	resp.UserID = row.ID
	resp.Token = token
	resp.ExpiresIn = expiresIn

	return resp, nil
}
