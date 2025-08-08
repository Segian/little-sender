package api

import (
	"github.com/gin-gonic/gin"
)

func NewSenderRouter(context *gin.Context) {

	context.JSON(200, gin.H{
		"message": "Sender router is not yet implemented",
	})
}
