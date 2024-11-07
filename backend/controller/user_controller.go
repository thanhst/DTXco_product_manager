package controller

import (
	"encoding/json"
	"net/http"
	"product_manage/model"
	"product_manage/service"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	err := uc.service.Register(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var creds model.User
	json.NewDecoder(r.Body).Decode(&creds)
	token, err := uc.service.Login(creds.Username, creds.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
