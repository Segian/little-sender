package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api")
	endpoint := api.Group("/endpoint")
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
	execution := router.Group("/execution")
	{
		execution.GET("/:id", ExecuteOneEndpoint)
	}
	return router
}
