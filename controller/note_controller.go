package controller

import (
	"cek/auth"

	"github.com/gin-gonic/gin"
)

type NoteController interface {
	FindAll(c *gin.Context, auth *auth.AccessDetails)
	Create(c *gin.Context, auth *auth.AccessDetails)
	Update(c *gin.Context, auth *auth.AccessDetails)
	Delete(c *gin.Context, auth *auth.AccessDetails)
}
