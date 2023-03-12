package controller

import (
	"cek/helper"
	"cek/service"

	"github.com/gin-gonic/gin"
)

type FileControllerImpl struct {
	FileService service.FileService
}

func NewFileController(fileService service.FileService) FileController {
	return &FileControllerImpl{
		FileService: fileService,
	}
}

func (controller *FileControllerImpl) FindFileDepletion(c *gin.Context) {
	file := c.Param("file")
	fileName := helper.PathProfile + file
	fileResponse := controller.FileService.FindFileDepletion(fileName)

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Write(fileResponse)
}
