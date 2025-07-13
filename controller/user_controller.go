package controller

import "github.com/example/user-service/service"

// UserController orchestrates user use cases.
type UserController struct {
	svc *service.UserService
}

func NewUserController(svc *service.UserService) *UserController {
	return &UserController{svc: svc}
}
