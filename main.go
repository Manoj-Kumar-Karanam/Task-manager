package main

import (
	"log"
	"task_manager/config"
	"task_manager/models"

)

func main() {
	config.ConnectDatabase();
	err := config.DB.AutoMigrate(&models.User{}, &models.Task{})
    if err != nil {
        log.Fatal("Failed to auto migrate models:", err)
    }

}

