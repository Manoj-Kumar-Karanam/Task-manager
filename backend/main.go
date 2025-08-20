package main

import (
	// "fmt"
	// "fmt"
	"log"
	"task_manager/config"
	// "task_manager/controllers"
	"task_manager/models"
	"time"
	"task_manager/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase();
	err := config.DB.AutoMigrate(&models.User{}, &models.Task{})
    if err != nil {
        log.Fatal("Failed to auto migrate models:", err)
    }
	r := gin.Default()
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // or "*" to allow all
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))
	routers.TaskRoutes(r)
	routers.UserRoutes(r)
	r.Run(":8080")
}

