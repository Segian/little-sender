package routes

import (
	"net/http"

	"github.com/Segian/little-sender/internal/model/dto"
	"github.com/Segian/little-sender/internal/service"
	"github.com/Segian/little-sender/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func sendErrorIfMethodInvalid(context *gin.Context, method string) bool {
	if !util.ValidateMethod(method) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid HTTP method"})
		return true
	}
	return false
}

func NewEndpoint(context *gin.Context) {
	var endpoint dto.EndpointDto

	if err := context.ShouldBindJSON(&endpoint); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if sendErrorIfMethodInvalid(context, endpoint.Method) {
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

func GetEndpoints(context *gin.Context) {
	endpoints, err := service.GetEndpoints()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, endpoints)
}

func GetEndpointByID(context *gin.Context) {
	id := context.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	endpoint, err := service.GetEndpointByID(parsedID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, endpoint)
}

func UpdateEndpoint(context *gin.Context) {
	id := context.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var endpointDto dto.EndpointDtoUpdate
	if err := context.ShouldBindJSON(&endpointDto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if endpointDto.Method != "" && sendErrorIfMethodInvalid(context, endpointDto.Method) {
		return
	}

	endpoint, err := service.UpdateEndpoint(parsedID, endpointDto)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, endpoint)
}

func DeleteEndpoint(context *gin.Context) {
	id := context.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := service.DeleteEndpoint(parsedID); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Endpoint deleted successfully"})
}
