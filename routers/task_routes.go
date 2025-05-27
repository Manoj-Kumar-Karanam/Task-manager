package routers

import (
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
)

func TaskRoutes(r *gin.Engine) {
	taskGroup := r.Group("/tasks") 
	{
		taskGroup.GET("/", controllers.GetallTasks)
		taskGroup.GET("/:id", controllers.GetTaskbyId)
		taskGroup.POST("/", controllers.CreateTask)
		taskGroup.PUT("/:id", controllers.UpdateTask)
		taskGroup.DELETE("/:id", controllers.DeleteTask)
	}
	
}

