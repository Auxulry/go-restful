// Package handler describe http handler
package handler

import (
	"net/http"

	"github.com/MochamadAkbar/go-restful/api"
	"github.com/MochamadAkbar/go-restful/common"
	"github.com/MochamadAkbar/go-restful/entity"
	"github.com/MochamadAkbar/go-restful/usecase"
	"github.com/go-chi/chi/v5"
)

type IUserHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	Usecase usecase.IUserUsecase
}

func NewUserHandler(usecase usecase.IUserUsecase, router *chi.Mux) http.Handler {
	handler := &UserHandler{
		Usecase: usecase,
	}

	router.Post("/api/register", handler.Register)
	router.Post("/api/login", handler.Login)

	return router
}

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req api.UserRequest
	res := api.Response{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
	}

	err := common.SerializeRequest(r, &req)
	if err != nil {
		panic(err)
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	resp, err := handler.Usecase.Register(r.Context(), user)
	if err != nil {
		panic(err)
	}

	res.Data = resp

	err = common.SerializeWriter(w, res.Code, res)
	if err != nil {
		panic(err)
	}
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req api.UserRequest
	res := api.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	err := common.SerializeRequest(r, &req)
	if err != nil {
		panic(err)
	}
	user := &entity.User{
		Email:    req.Email,
		Password: req.Password,
	}
	resp, err := handler.Usecase.Login(r.Context(), user)
	if err != nil {
		panic(err)
	}
	res.Data = resp
	err = common.SerializeWriter(w, res.Code, res)
	if err != nil {
		panic(err)
	}
}
