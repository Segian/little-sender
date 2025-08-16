package routes

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	endpoint := router.Group("/endpoint")
	{
		endpoint.POST("", NewEndpoint)
		endpoint.GET("", GetEndpoints)
		endpoint.GET("/:id", GetEndpointByID)
		endpoint.PUT("/:id", UpdateEndpoint)
		endpoint.DELETE("/:id", DeleteEndpoint)
	}
	health := router.Group("/health")
	{
		health.GET("", HealthCheck)
	}
	return router
}
