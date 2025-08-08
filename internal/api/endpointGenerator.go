package api

import (
	"net/http"

	"github.com/Segian/little-sender/internal/model"
	"github.com/Segian/little-sender/internal/service"
	"github.com/gin-gonic/gin"
)

func NewEndpoint(context *gin.Context) {
	var endpointModel model.EndpointModel

	if err := context.ShouldBindJSON(&endpointModel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	endpointService := service.NewEndpointService(endpointModel)

	context.JSON(http.StatusOK, gin.H{
		"dataSended": endpointService,
		"message":    "Endpoint created successfully",
	})
}
