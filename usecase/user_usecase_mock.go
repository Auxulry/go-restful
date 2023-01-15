package usecase

import (
	"context"

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
		return api.UserResponse{}, args.Error(1)
	} else {
		result := args.Get(0).(api.UserResponse)
		return result, nil
	}
}

func (usecase *UserUsecaseMock) Login(_ context.Context, user *entity.User) (api.UserResponse, error) {
	args := usecase.Mock.Called(user)

	if args.Get(0) == nil {
		return api.UserResponse{}, args.Error(1)
	} else {
		result := args.Get(0).(api.UserResponse)
		return result, nil
	}
}
