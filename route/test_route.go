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

func TestRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	testService := service.NewTestService(
		repository.NewTestRepository(),
		db,
		validate,
	)

	TestRoute := controller.NewTestController(testService)
	router.GET("/test", auth.Auth(TestRoute.FindAll, []string{auth.RoleAdministrator}))

	// router.GET("/test",(TestRoute.FindAll, []string{}))
}
