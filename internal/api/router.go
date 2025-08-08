package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	endpoint := router.Group("/endpoint")
	{
		endpoint.POST("", NewEndpoint)
	}
	health := router.Group("/health")
	{
		health.GET("", HealthCheck)
	}
	return router
}
