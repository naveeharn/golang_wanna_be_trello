package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naveeharn/golang_wanna_be_trello/dto"
	"github.com/naveeharn/golang_wanna_be_trello/helper"
	"github.com/naveeharn/golang_wanna_be_trello/service"
)

type DashboardController interface {
	CreateDashboard(ctx *gin.Context)
	UpdateDashboard(ctx *gin.Context)
	// CreateNote(ctx *gin.Context)
}

type dashboardController struct {
	dashboardService service.DashboardService
}

func NewDashboardController(dashboardService service.DashboardService) DashboardController {
	return &dashboardController{
		dashboardService: dashboardService,
	}
}

func (controller *dashboardController) CreateDashboard(ctx *gin.Context) {
	dashboardBeforeCreate := dto.DashboardCreateDTO{}
	if err := ctx.ShouldBind(&dashboardBeforeCreate); err != nil {
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

	teamId := ctx.Param("teamId")
	if teamId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "teamId doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	dashboardBeforeCreate.OwnerUserId = userId.(string)
	dashboardBeforeCreate.TeamId = teamId
	updatedTeam, err := controller.dashboardService.CreateDashboard(dashboardBeforeCreate)
	if err != nil {
		response := helper.CreateErrorResponse("Failed to process request", "Failed to create new dashboard in team", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.CreateResponse(true, "Create dashboard response complete", updatedTeam)
	ctx.JSON(http.StatusCreated, response)
}

func (controller *dashboardController) UpdateDashboard(ctx *gin.Context) {
	dashboardBeforeUpdate := dto.DashboardNameUpdateDTO{}
	if err := ctx.ShouldBind(&dashboardBeforeUpdate); err != nil {
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

	teamId := ctx.Param("teamId")
	if teamId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "teamId doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	dashboardId := ctx.Param("dashboardId")
	if dashboardId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "dashboardId doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}
	dashboardBeforeUpdate.Id = dashboardId
	dashboardBeforeUpdate.OwnerUserId = userId.(string)
	dashboardBeforeUpdate.TeamId = teamId
	updatedTeam, err := controller.dashboardService.UpdateDashboard(dashboardBeforeUpdate)
	if err != nil {
		response := helper.CreateErrorResponse("Failed to process request", "Failed to update dashboard", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.CreateResponse(true, "Update dashboard response complete", updatedTeam)
	ctx.JSON(http.StatusCreated, response)
}
