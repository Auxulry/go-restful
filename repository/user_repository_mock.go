package repository

import (
	"context"

	"github.com/MochamadAkbar/go-restful/entity"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) Register(_ context.Context, user *entity.User) (entity.User, error) {
	args := repository.Mock.Called(user)

	if args.Get(0) == nil {
		return entity.User{}, args.Error(1)
	} else {
		result := args.Get(0).(entity.User)
		return result, nil
	}
}

func (repository *UserRepositoryMock) Login(ctx context.Context, user *entity.User) (entity.User, error) {
	args := repository.Mock.Called(user)

	if args.Get(0) == nil {
		return entity.User{}, args.Error(1)
	} else {
		result := args.Get(0).(entity.User)
		return result, nil
	}
}
