package controller

import (
	"cek/auth"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	// CRUD
	Create(context *gin.Context, auth *auth.AccessDetails)
	Delete(context *gin.Context, auth *auth.AccessDetails)
	Update(context *gin.Context, auth *auth.AccessDetails)
	FindAll(context *gin.Context, auth *auth.AccessDetails)
	ResetPassword(context *gin.Context, auth *auth.AccessDetails)
	ChangePassword(context *gin.Context, auth *auth.AccessDetails)

	// Auth
	Login(context *gin.Context)
	RefreshToken(context *gin.Context)
	Registration(context *gin.Context)
}
