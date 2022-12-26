// Package handler describe http handler
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MochamadAkbar/go-restful/api"
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

func InitHandler(usecase usecase.IUserUsecase, router *chi.Mux) IUserHandler {
	handler := &UserHandler{
		Usecase: usecase,
	}

	router.Post("/api/register", handler.Register)
	router.Post("/api/login", handler.Login)

	return handler
}

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req api.UserRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
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
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(resp)
	if err != nil {
		panic(err)
	}
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req api.UserRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
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
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(resp)
	if err != nil {
		panic(err)
	}
}
