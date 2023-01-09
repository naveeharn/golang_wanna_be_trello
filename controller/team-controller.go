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
	GetTeamById(ctx *gin.Context)
	GetTeamsByOwnerUserId(ctx *gin.Context)
	AddMember(ctx *gin.Context)
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

func (controller *teamController) GetTeamById(ctx *gin.Context) {
	teamId := ctx.Param("id")
	if teamId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "Team id doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}
	userId, ok := ctx.Get("userId")
	if !ok || userId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "User id fron JWT token doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	team, err := controller.teamService.GetTeamById(teamId, userId.(string))
	if err != nil {
		response := helper.CreateErrorResponse("Failed to process request", "Failed to get team by id", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.CreateResponse(true, "Create team response complete", team)
	ctx.JSON(http.StatusCreated, response)
}

func (controller *teamController) GetTeamsByOwnerUserId(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok || userId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "User id fron JWT token doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	teams, err := controller.teamService.GetTeamsByOwnerUserId(userId.(string))
	if err != nil {
		response := helper.CreateErrorResponse("Failed to process request", "Failed to get team by owner user id", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.CreateResponse(true, "Get Teams by owner user id team response complete", teams)
	ctx.JSON(http.StatusCreated, response)
}

func (controller *teamController) AddMember(ctx *gin.Context) {
	newMember := dto.TeamAddMemberDTO{}
	if err := ctx.ShouldBind(&newMember); err != nil {
		response := helper.CreateErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	teamId := ctx.Param("id")
	if teamId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "Team id doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}
	userId, ok := ctx.Get("userId")
	if !ok || userId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "User id fron JWT token doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	team, err := controller.teamService.AddMember(teamId, userId.(string), newMember.Email)
	if err != nil {
		response := helper.CreateErrorResponse("Failed to process request", "Failed to add new member in team", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.CreateResponse(true, "Create team response complete", team)
	ctx.JSON(http.StatusCreated, response)
}
