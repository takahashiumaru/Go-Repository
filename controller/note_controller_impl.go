package controller

import (
	"net/http"
	"strconv"

	"cek/auth"
	"cek/helper"
	"cek/model/web"
	"cek/service"

	"github.com/gin-gonic/gin"
)

type NoteControllerImpl struct {
	NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) NoteController {
	return &NoteControllerImpl{
		NoteService: noteService,
	}
}

func (controller *NoteControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "subject.eq")
	noteResponses := controller.NoteService.FindAll(auth, &filters)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(noteResponses),
		Data:    noteResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *NoteControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {

	request := web.NoteCreateRequest{}
	helper.ReadFromRequestBody(c, &request)
	noteResponse := controller.NoteService.Create(auth, &request)

	webResponse := web.WebResponse{
		Success: true,
		Message: "Note Created Successfully",
		Data:    noteResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *NoteControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	note := c.Param("id")
	noteID, err := strconv.Atoi(note)
	helper.PanicIfError(err)
	request := web.NoteUpdateRequest{}
	helper.ReadFromRequestBody(c, &request)
	noteResponse := controller.NoteService.Update(auth, &noteID, &request)

	webResponse := web.WebResponse{
		Success: true,
		Message: "Note Updated Successfully",
		Data:    noteResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *NoteControllerImpl) Delete(c *gin.Context, auth *auth.AccessDetails) {
	note := c.Param("id")
	noteID, err := strconv.Atoi(note)
	helper.PanicIfError(err)

	controller.NoteService.Delete(auth, &noteID)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Event Detail deleted successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}
