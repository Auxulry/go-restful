// Package handler describe http handler
package handler

import (
	"net/http"

	"github.com/MochamadAkbar/go-restful/api"
	commonErr "github.com/MochamadAkbar/go-restful/common/errors"
	"github.com/MochamadAkbar/go-restful/common/serializer"
	"github.com/MochamadAkbar/go-restful/entity"
	"github.com/MochamadAkbar/go-restful/usecase"
	"github.com/go-chi/chi/v5"
)

type UserHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type UserHandlerImpl struct {
	Usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase, router chi.Router) error {
	handler := &UserHandlerImpl{
		Usecase: usecase,
	}

	router.Post("/authentication/register", handler.Register)
	router.Post("/authentication/login", handler.Login)

	return nil
}

func (handler *UserHandlerImpl) Register(w http.ResponseWriter, r *http.Request) {
	var req api.UserRequest

	serializer.SerializeRequest(r, &req)

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	defer serializer.SerializeCloseRequest(r)

	resp, err := handler.Usecase.Register(r.Context(), user)
	if err != nil {
		result := api.ErrResponse{
			Code:    commonErr.GetHTTPCode(err),
			Status:  http.StatusText(commonErr.GetHTTPCode(err)),
			Message: err.Error(),
		}

		serializer.SerializeWriter(w, result.Code, result)
		return
	}

	result := api.Response{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   resp,
	}

	serializer.SerializeWriter(w, result.Code, result)
}

func (handler *UserHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var req api.UserRequest

	serializer.SerializeRequest(r, &req)

	user := &entity.User{
		Email:    req.Email,
		Password: req.Password,
	}

	defer serializer.SerializeCloseRequest(r)

	resp, err := handler.Usecase.Login(r.Context(), user)
	if err != nil {
		result := api.ErrResponse{
			Code:    commonErr.GetHTTPCode(err),
			Status:  http.StatusText(commonErr.GetHTTPCode(err)),
			Message: err.Error(),
		}

		serializer.SerializeWriter(w, result.Code, result)
		return
	}

	result := api.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   resp,
	}

	serializer.SerializeWriter(w, result.Code, result)
}
