package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Define routes
	router.POST("/properties", CreateProperty)
	router.GET("/properties", GetProperties)
	router.HEAD("/properties", GetProperties)

	// Health check
	router.GET("/health", HealthCheck)
	router.HEAD("/health", HealthCheck)
}
