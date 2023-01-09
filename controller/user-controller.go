package controller

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/naveeharn/golang_wanna_be_trello/dto"
	"github.com/naveeharn/golang_wanna_be_trello/helper"
	"github.com/naveeharn/golang_wanna_be_trello/service"
)

type UserController interface {
	UpdateUser(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	Profile(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
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
	userId := ctx.Param("id")
	if userId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "User id doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	user, err := controller.userService.GetUserById(userId)
	if err != nil {
		response := helper.CreateErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.CreateResponse(true, "Get user by id response complete", user)
	ctx.JSON(http.StatusOK, response)
}

func (controller *userController) Profile(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok || userId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "User id fron JWT token doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}
	helper.LoggerErrorPath(runtime.Caller(0))

	user, err := controller.userService.GetUserById(userId.(string))
	if err != nil {
		response := helper.CreateErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	helper.LoggerErrorPath(runtime.Caller(0))

	response := helper.CreateResponse(true, "Get profile user response complete", user)
	ctx.JSON(http.StatusOK, response)
}

func (controller *userController) UpdateUser(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok || userId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "User id fron JWT token doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	userUpdateDTO := dto.UserUpdateDTO{}
	err := ctx.ShouldBind(&userUpdateDTO)
	if err != nil {
		response := helper.CreateErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	userUpdateDTO.Id = userId.(string)

	// userBeforeUpdate, err := controller.userService.GetUserById(userUpdateDTO.Id)
	// if err != nil {
	// 	response := helper.CreateErrorResponse("userId from accessToken does not exists", err.Error(), nil)
	// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	// 	return
	// }

	updatedUser, err := controller.userService.Update(userUpdateDTO)
	if err != nil {
		response := helper.CreateErrorResponse("userId from accessToken does not exists", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	response := helper.CreateResponse(true, "User update response complete", updatedUser)
	ctx.JSON(http.StatusOK, response)
}

func (controller *userController) ResetPassword(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok || userId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "User id fron JWT token doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	userResetPasswordDTO := dto.UserResetPasswordDTO{}
	err := ctx.ShouldBind(&userResetPasswordDTO)
	if err != nil {
		response := helper.CreateErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userResetPasswordDTO.Id = userId.(string)
	updatedUser, err := controller.userService.ResetPassword(userResetPasswordDTO)
	if err != nil {
		response := helper.CreateErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.CreateResponse(true, "Reset password updatedUser response complete", updatedUser)
	ctx.JSON(http.StatusOK, response)
}
