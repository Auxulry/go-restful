package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MochamadAkbar/go-restful/config"
	"github.com/MochamadAkbar/go-restful/entity"
	"github.com/MochamadAkbar/go-restful/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userHandler = UserHandler{Usecase: userUsecase}
	userUsecase = &usecase.UserUsecaseMock{Mock: mock.Mock{}}
)

func TestUserHandler_Login(t *testing.T) {
	t.Run("Test User Handler Login Success", func(t *testing.T) {
		user := new(entity.User)
		user.Email = "test@gmail.com"
		user.Password = "password"

		router := config.NewRouter()
		router.Post("/api/login", userHandler.Login)

		requestBody := strings.NewReader(`{ "email": "test@gmail.com", "password": "password" }`)
		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/login", requestBody)

		recorder := httptest.NewRecorder()
		userUsecase.Mock.On("Login", user).Return(true)

		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close() //nolint:errcheck
		assert.Equal(t, "200 OK", response.Status)
		assert.Equal(t, http.StatusOK, response.StatusCode)
	})
}

func TestUserHandler_Register(t *testing.T) {
	t.Run("Test User Handler Register Success", func(t *testing.T) {
		user := new(entity.User)
		user.Name = "test"
		user.Email = "test@gmail.com"
		user.Password = "password"

		router := config.NewRouter()
		router.Post("/api/register", userHandler.Register)

		requestBody := strings.NewReader(`{ "name": "test", "email": "test@gmail.com", "password": "password" }`)
		request := httptest.NewRequest(http.MethodPost, "http://localhost:5000/api/register", requestBody)

		recorder := httptest.NewRecorder()
		userUsecase.Mock.On("Register", user).Return(true)

		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close() //nolint:errcheck
		assert.Equal(t, "201 Created", response.Status)
		assert.Equal(t, http.StatusCreated, response.StatusCode)
	})
}
