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

func UserRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {
	userRepository := repository.NewUserRepository()
	sessionRepository := repository.NewSessionRepository()
	// marketingStructureRepository := repository.NewMarketingStructureRepository()

	userService := service.NewUserService(userRepository, sessionRepository, db, validate)
	userController := controller.NewUserController(userService)

	router.DELETE("/users/:id", auth.Auth(userController.Delete, []string{auth.RoleAdministrator}))
	router.GET("/users", auth.Auth(userController.FindAll, []string{auth.RoleAdministrator}))
	router.POST("/users", auth.Auth(userController.Create, []string{auth.RoleAdministrator}))
	router.PUT("/users/:id", auth.Auth(userController.Update, []string{auth.RoleAdministrator}))
	router.PUT("/users-change-password/:id", auth.Auth(userController.ChangePassword, []string{}))
	router.PUT("/users-reset-password/:id", auth.Auth(userController.ResetPassword, []string{}))

	router.GET("/users/login", userController.Login)
	router.POST("/users/refresh-token", userController.RefreshToken)
	router.POST("/users/registration", userController.Registration)
}
