package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"task_manager/config"
	"task_manager/models"
	"errors"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)


func GetallTasks(c* gin.Context) {
	var tasks []models.Task
	err := config.DB.Find(&tasks).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "failed to retrieve tasks"})
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskbyId(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	fmt.Println("User with id",userID,"with task id ", idUint)
	var task models.Task
	if err := config.DB.
		Where("id = ? AND user_id = ?", idUint, userID).
		First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	userId, exists := c.Get("user_id")
	 if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "User Unauthorized",
		})
		return
	 }

	var newtask models.Task

	// Try binding JSON to struct
	if err := c.ShouldBindJSON(&newtask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input: " + err.Error(),
		})
		return
	}
	newtask.UserID = userId.(uint)
	// Try saving to DB
	if newtask.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "title Cannot be empty",
		})
	}
	if err := config.DB.Create(&newtask).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error: " + err.Error(),
		})
		return
	}

	// Success
	c.JSON(http.StatusCreated, newtask)
}


func DeleteTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	var task models.Task
	if err := config.DB.First(&task, idUint).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if task.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this task"})
		return
	}

	if err := config.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}


func GetUsertasks(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "user not authorized",
		})
		return
	}
	var tasks []models.Task 
	if err := config.DB.Where("user_id=?", userId).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "cannot fetch tasks",
		})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	if err := config.DB.First(&task, idUint).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if task.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this task"})
		return
	}

	//new task struct that needs to be updated.
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.Title = input.Title
	task.Description = input.Description
	task.Completed = input.Completed

	if err := config.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully", "task": task})
}



func GetTaskDetails(c *gin.Context) {
    taskID := c.Param("task_id")
    var task models.Task

    db := config.DB.Where("id = ?", taskID).First(&task)
    if db.Error != nil {
        if errors.Is(db.Error, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch details"})
        }
        return
    }

    c.JSON(http.StatusOK, task)
}
