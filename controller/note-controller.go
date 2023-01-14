package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naveeharn/golang_wanna_be_trello/dto"
	"github.com/naveeharn/golang_wanna_be_trello/helper"
	"github.com/naveeharn/golang_wanna_be_trello/service"
)

type NoteController interface {
	CreateNote(ctx *gin.Context)
	UpdateNote(ctx *gin.Context)
	DeleteNote(ctx *gin.Context)
}

type noteController struct {
	noteService service.NoteService
}

func NewNoteController(nodeService service.NoteService) NoteController {
	return &noteController{
		noteService: nodeService,
	}
}

func (controller *noteController) CreateNote(ctx *gin.Context) {
	noteBeforeCreate := dto.NoteCreateDTO{}
	if err := ctx.ShouldBind(&noteBeforeCreate); err != nil {
		response := helper.CreateErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId, ok := ctx.Get("userId")
	if !ok || userId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "User id from JWT Token doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}
	dashboardId := ctx.Param("dashboardId")
	if dashboardId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "dashboardId doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	noteBeforeCreate.OwnerUserId = userId.(string)
	noteBeforeCreate.DashboardId = dashboardId
	updatedDashboard, err := controller.noteService.CreateNote(noteBeforeCreate)
	if err != nil {
		response := helper.CreateErrorResponse("Failed to create new note by dashboard id", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.CreateResponse(true, "Create note response complete", updatedDashboard)
	ctx.JSON(http.StatusCreated, response)
}

func (controller *noteController) DeleteNote(ctx *gin.Context) {

}

func (controller *noteController) UpdateNote(ctx *gin.Context) {

}
