package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naveeharn/golang_wanna_be_trello/dto"
	"github.com/naveeharn/golang_wanna_be_trello/helper"
	"github.com/naveeharn/golang_wanna_be_trello/service"
)

type TeamController interface {
	CreateTeam(ctx *gin.Context)
}

type teamController struct {
	teamService service.TeamService
}

func NewTeamController(teamService service.TeamService) TeamController {
	return &teamController{
		teamService: teamService,
	}
}

func (controller *teamController) CreateTeam(ctx *gin.Context) {
	teamBeforeCreate := dto.TeamCreateDTO{}
	if err := ctx.ShouldBind(&teamBeforeCreate); err != nil {
		response := helper.CreateErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId, ok := ctx.Get("userId")
	if !ok || userId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "User id fron JWT token doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	teamBeforeCreate.OwnerUserId = userId.(string)
	createdTeam, err := controller.teamService.CreateTeam(teamBeforeCreate)
	if err != nil {
		response := helper.CreateErrorResponse("Failed to process request", "Failed to create new team", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.CreateResponse(true, "Create team response complete", createdTeam)
	ctx.JSON(http.StatusCreated, response)
}
