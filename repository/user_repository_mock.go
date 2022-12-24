package repository

import (
	"context"

	"github.com/MochamadAkbar/go-restful/entity"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) Register(ctx context.Context, user *entity.User) bool {
	args := repository.Mock.Called(ctx, user)

	if args.Get(0) == nil {
		return false
	} else {
		return true
	}
}

func (repository *UserRepositoryMock) Login(ctx context.Context, user *entity.User) (string, bool) {
	args := repository.Mock.Called(ctx, user)

	if args.Get(0) == nil {
		return "", false
	} else {
		return "UUID", true
	}
}
