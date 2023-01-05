package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/naveeharn/golang_wanna_be_trello/service"
)

type UserController interface {
	Update(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (controller *userController) GetUserById(ctx *gin.Context) {
	panic("unimplemented")
}

func (controller *userController) Update(ctx *gin.Context) {
	panic("unimplemented")
}
