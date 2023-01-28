package controller

import (
	"cek/auth"

	"github.com/gin-gonic/gin"
)

type TestController interface {
	FindAll(c *gin.Context, auth *auth.AccessDetails)
}
