package route

import (
	"cek/auth"
	"cek/controller"
	"cek/repository"
	"cek/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func NoteRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	noteService := service.NewNoteService(
		repository.NewNoteRepository(),
		db,
		validate,
	)

	NoteRoute := controller.NewNoteController(noteService)
	router.GET("/note", auth.Auth(NoteRoute.FindAll, []string{}))
	router.POST("/note", auth.Auth(NoteRoute.Create, []string{}))
	router.PUT("/note/:id", auth.Auth(NoteRoute.Update, []string{}))
	router.DELETE("/note/:id", auth.Auth(NoteRoute.Delete, []string{}))
}
