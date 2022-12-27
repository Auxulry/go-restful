package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/MochamadAkbar/go-restful/api"
	"github.com/MochamadAkbar/go-restful/entity"
	"github.com/MochamadAkbar/go-restful/repository"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	Mock       mock.Mock
	Repository repository.UserRepositoryMock
}

func (usecase *UserUsecaseMock) Register(_ context.Context, user *entity.User) (api.UserResponse, error) {
	args := usecase.Mock.Called(user)
	if args.Get(0) == nil {
		return api.UserResponse{}, errors.New("failed to register")
	} else {
		return api.UserResponse{
			UserID:    "userid",
			Token:     "token",
			ExpiresIn: time.Now().Add(time.Duration(1) * time.Minute).Unix(),
		}, nil
	}
}

func (usecase *UserUsecaseMock) Login(_ context.Context, user *entity.User) (api.UserResponse, error) {
	args := usecase.Mock.Called(user)

	if args.Get(0) == nil {
		return api.UserResponse{}, errors.New("failed to login")
	} else {
		usecase.Repository = repository.UserRepositoryMock{
			Mock: mock.Mock{},
		}

		return api.UserResponse{
			UserID:    "userid",
			Token:     "token",
			ExpiresIn: time.Now().Add(time.Duration(1) * time.Minute).Unix(),
		}, nil
	}
}
