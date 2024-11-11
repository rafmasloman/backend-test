package routes

import (
	"backend/api/src/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	
	customerRoutes := router.Group("/api/v1/customer") 
	
	{
		customerRoutes.GET("/", controllers.GetAllCustomers)
		customerRoutes.POST("/", controllers.CreateCustomer)
		customerRoutes.GET("/:id", controllers.GetDetailCustomer)
		customerRoutes.PUT("/:id", controllers.UpdateCustomer)
		customerRoutes.DELETE("/:id", controllers.DeleteCustomer)
	}

	return router
}