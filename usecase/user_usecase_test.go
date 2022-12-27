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

var (
	userUsecase        = UserUsecase{Repository: userRepositoryMock}
	userRepositoryMock = &repository.UserRepositoryMock{Mock: mock.Mock{}}
)

func TestUserUsecase_Register(t *testing.T) {
	t.Run("Test User Use Case Register Success", func(t *testing.T) {
		var resp api.UserResponse

		ctx := context.Background()

		user := new(entity.User)

		user.Name = "John Doe"
		user.Email = "johndoe@gmail.com"

		userRepositoryMock.Mock.On("Register", user).Return(true)
		userRepositoryMock.Mock.On("Login", user).Return("UUID", true)

		result, err := userUsecase.Register(ctx, user)

		resp.UserID = result.UserID
		resp.Token = result.Token

		assert.Nil(t, err)
		assert.Equal(t, "UUID", resp.UserID)
		assert.NotEqual(t, "", resp.Token)
	})

	t.Run("Test User Use Case Register Failed", func(t *testing.T) {
		var resp api.UserResponse
		ctx := context.Background()

		user := new(entity.User)

		userRepositoryMock.Mock.On("Register", user).Return(nil)
		userRepositoryMock.Mock.On("Login", user).Return(nil)

		result, err := userUsecase.Register(ctx, user)

		assert.NotNil(t, err)
		assert.Equal(t, resp.UserID, result.UserID)
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

		userRepositoryMock.Mock.On("Login", user).Return("UUID", true)

		result, err := userUsecase.Login(ctx, user)

		resp.UserID = result.UserID
		resp.Token = result.Token

		assert.Nil(t, err)
		assert.Equal(t, "UUID", resp.UserID)
		assert.NotEqual(t, "", resp.Token)
	})

	t.Run("Test User use case Login Fail", func(t *testing.T) {
		var resp api.UserResponse
		ctx := context.Background()

		user := new(entity.User)

		userRepositoryMock.Mock.On("Login", user).Return(nil)

		result, err := userUsecase.Login(ctx, user)

		assert.NotNil(t, err)
		assert.Equal(t, resp.UserID, result.UserID)
		assert.Equal(t, resp.Token, result.Token)
	})
}
