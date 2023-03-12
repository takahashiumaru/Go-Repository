package route

import (
	"cek/controller"
	"cek/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func FileRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	fileService := service.NewFileService()

	fileController := controller.NewFileController(fileService)

	router.GET("/file/profile/:file", fileController.FindFileDepletion)
}
