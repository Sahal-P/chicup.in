package router

import (
	"chicup/internal/handlers"
	"chicup/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {
	r := gin.Default()

	// Apply middleware
	r.Use(middleware.CORS())

	// Define routes
	api := r.Group("/api")
	{
		api.POST("/register", userHandler.RegisterHandler)
		// Add other routes here
	}

	return r
}
