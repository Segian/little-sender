package api

import (
	"net/http"

	"github.com/Segian/little-sender/internal/dto"
	"github.com/Segian/little-sender/internal/service"
	"github.com/gin-gonic/gin"
)

func NewEndpoint(context *gin.Context) {
	var endpoint dto.EndpointDto

	if err := context.ShouldBindJSON(&endpoint); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	endpointService, err := service.NewEndpointService(endpoint)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create endpoint"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"dataSended": endpointService,
		"message":    "Endpoint created successfully",
	})
}
