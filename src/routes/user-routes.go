package routes

import (
	"backend/api/src/controllers"
	"backend/api/src/middleware"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {

	router := gin.Default()

	router.Use(middleware.CorsMiddleware())	

	
	userRoutes := router.Group("/users") 
	{
		userRoutes.GET("/", controllers.GetAllUsers)
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/:id", controllers.GetDetailUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	return router
}