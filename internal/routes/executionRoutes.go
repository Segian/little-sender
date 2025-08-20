package routes

import (
	"github.com/Segian/little-sender/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ExecuteOneEndpoint(context *gin.Context) {
	id := context.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	executionService, err := service.ExecuteEndpoint(parsedID)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to execute endpoint"})
		return
	}
	context.JSON(200, executionService)
}
