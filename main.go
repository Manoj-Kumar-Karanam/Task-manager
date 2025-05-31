package main

import (
	// "fmt"
	// "fmt"
	"log"
	"task_manager/config"
	// "task_manager/controllers"
	"task_manager/models"
	"task_manager/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase();
	err := config.DB.AutoMigrate(&models.User{}, &models.Task{})
    if err != nil {
        log.Fatal("Failed to auto migrate models:", err)
    }
	r := gin.Default()
	routers.TaskRoutes(r)
	routers.UserRoutes(r)
	r.Run(":8080")
}

