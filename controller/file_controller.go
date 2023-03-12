package controller

import (
	"github.com/gin-gonic/gin"
)

type FileController interface {
	FindFileDepletion(context *gin.Context)
}
