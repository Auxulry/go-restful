package usecase

import (
	"context"
	"testing"

	"github.com/MochamadAkbar/go-restful/api"
	"github.com/MochamadAkbar/go-restful/entity"
	"github.com/MochamadAkbar/go-restful/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepositoryMock = &repository.UserRepositoryMock{Mock: mock.Mock{}}

var userUsecase = UserUsecase{Repository: userRepositoryMock}

func TestUserUsecase_Register(t *testing.T) {
	t.Run("Test User Use Case Register Success", func(t *testing.T) {
		var resp api.UserResponse

		ctx := context.Background()

		user := new(entity.User)

		user.Name = "John Doe"
		user.Email = "johndoe@gmail.com"

		userRepositoryMock.Mock.On("Register", ctx, user).Return(true)
		userRepositoryMock.Mock.On("Login", ctx, user).Return("UUID", true)

		result, err := userUsecase.Register(ctx, user)

		resp.UUID = result.UUID
		resp.Token = result.Token

		assert.Nil(t, err)
		assert.Equal(t, "UUID", resp.UUID)
		assert.Equal(t, "Token", resp.Token)
	})

	t.Run("Test User Use Case Register Failed", func(t *testing.T) {
		var resp api.UserResponse
		ctx := context.Background()

		user := new(entity.User)

		userRepositoryMock.Mock.On("Register", ctx, user).Return(nil)
		userRepositoryMock.Mock.On("Login", ctx, user).Return(nil)

		result, err := userUsecase.Register(ctx, user)

		assert.NotNil(t, err)
		assert.Equal(t, resp.UUID, result.UUID)
		assert.Equal(t, resp.Token, result.Token)
	})
}

func TestUserUsecase_Login(t *testing.T) {
	t.Run("Test User Use Case Login Success", func(t *testing.T) {
		var resp api.UserResponse

		ctx := context.Background()

		user := new(entity.User)

		user.Email = "johndoe@gmail.com"
		user.Password = "test"

		userRepositoryMock.Mock.On("Login", ctx, user).Return("UUID", true)

		result, err := userUsecase.Login(ctx, user)

		resp.UUID = result.UUID
		resp.Token = result.Token

		assert.Nil(t, err)
		assert.Equal(t, "UUID", resp.UUID)
		assert.Equal(t, "Token", resp.Token)
	})

	t.Run("Test User use case Login Fail", func(t *testing.T) {
		var resp api.UserResponse
		ctx := context.Background()

		user := new(entity.User)

		userRepositoryMock.Mock.On("Login", ctx, user).Return(nil)

		result, err := userUsecase.Login(ctx, user)

		assert.NotNil(t, err)
		assert.Equal(t, resp.UUID, result.UUID)
		assert.Equal(t, resp.Token, result.Token)
	})
}
