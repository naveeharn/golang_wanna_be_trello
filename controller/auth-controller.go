package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naveeharn/golang_wanna_be_trello/dto"
	"github.com/naveeharn/golang_wanna_be_trello/helper"
	"github.com/naveeharn/golang_wanna_be_trello/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (controller *authController) Login(ctx *gin.Context) {
	panic("unimplemented")
}

func (controller *authController) Register(ctx *gin.Context) {
	registerDTO := dto.UserCreateDTO{}

	if err := ctx.Bind(&registerDTO); err != nil {
		response := helper.CreateErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	log.Printf("%#v", registerDTO.BirthDate)

	if controller.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.CreateErrorResponse("Failed to process request", "Duplicated email", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	createdUser, err := controller.authService.CreateUser(registerDTO)
	if err != nil {
		response := helper.CreateErrorResponse("Failed to process request", "Failed to create new user", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	createdUser.AccessToken = controller.jwtService.GenerateToken(createdUser.Id)

	response := helper.CreateResponse(true, "Register response complete", createdUser)
	ctx.JSON(http.StatusCreated, response)
}
