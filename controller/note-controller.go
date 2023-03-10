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
		response := helper.CreateErrorResponse("Failed to process request", "Dashboard id from note doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	noteBeforeCreate.OwnerUserId = userId.(string)
	noteBeforeCreate.DashboardId = dashboardId
	updatedDashboard, err := controller.noteService.CreateNote(noteBeforeCreate)
	if err != nil {
		response := helper.CreateErrorResponse("Failed to create new note", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.CreateResponse(true, "Create note response complete", updatedDashboard)
	ctx.JSON(http.StatusCreated, response)
}

func (controller *noteController) UpdateNote(ctx *gin.Context) {
	noteBeforeUpdate := dto.NoteUpdateDTO{}
	if err := ctx.ShouldBind(&noteBeforeUpdate); err != nil {
		response := helper.CreateErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
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
		response := helper.CreateErrorResponse("Failed to process request", "Dashboard id from note doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}
	noteId := ctx.Param("noteId")
	if noteId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "Dashboard id from note doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	noteBeforeUpdate.Id = noteId
	noteBeforeUpdate.OwnerUserId = userId.(string)
	noteBeforeUpdate.DashboardId = dashboardId
	updatedDashboard, err := controller.noteService.UpdateNote(noteBeforeUpdate)
	if err != nil {
		response := helper.CreateErrorResponse("Failed to update note by id", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.CreateResponse(true, "Update note response complete", updatedDashboard)
	ctx.JSON(http.StatusCreated, response)
}

func (controller *noteController) DeleteNote(ctx *gin.Context) {
	noteBeforeDelete := dto.NoteDeleteDTO{}
	if err := ctx.ShouldBind(&noteBeforeDelete); err != nil {
		response := helper.CreateErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
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
		response := helper.CreateErrorResponse("Failed to process request", "Dashboard id from note doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}
	noteId := ctx.Param("noteId")
	if noteId == "" {
		response := helper.CreateErrorResponse("Failed to process request", "Dashboard id from note doesn't found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}
	noteBeforeDelete.Id = noteId
	noteBeforeDelete.DashboardId = dashboardId
	noteBeforeDelete.OwnerUserId = userId.(string)

	// log.Printf("\n\n%#v\n\n", noteBeforeDelete)
	updatedDashboard, err := controller.noteService.DeleteNote(noteBeforeDelete)
	if err != nil {
		response := helper.CreateErrorResponse("Failed to delete note by id", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.CreateResponse(true, "Update note response complete", updatedDashboard)
	ctx.JSON(http.StatusCreated, response)
}
