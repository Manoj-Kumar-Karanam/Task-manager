package routers

import (
	"github.com/gin-gonic/gin"
	"task_manager/middleware"
	"task_manager/controllers"
)

func TaskRoutes(router *gin.Engine) {
	
	router.GET("/tasks", controllers.GetallTasks)

	// Authenticated user routes
	auth := router.Group("/api")
	auth.Use(middleware.AuthMiddleware()) // This middleware sets userID into context
	{
		auth.POST("/tasks", controllers.CreateTask)         // Create a task
		auth.GET("/tasks", controllers.GetUsertasks)        // List only the user's tasks
		auth.GET("/tasks/:id", controllers.GetTaskbyId)     // Get one task (with ownership check)
		auth.PUT("/tasks/:id", controllers.UpdateTask)      // Update task (user only)
		auth.DELETE("/tasks/:id", controllers.DeleteTask)   // Delete task (user only)
	}
}