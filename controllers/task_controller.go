package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"task_manager/models"
	"task_manager/config"
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
	id := c.Param("id")
	var task models.Task
	err := config.DB.Find(&task, id)
	if err != nil {
    	c.JSON(http.StatusNotFound, gin.H{
        	"error": "Failed to retrieve task by id",
        	"id":    id,
    	})
    	return
	}
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var newtask models.Task

	// Try binding JSON to struct
	if err := c.ShouldBindJSON(&newtask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input: " + err.Error(),
		})
		return
	}

	// Try saving to DB
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
	id := c.Param("id")
	var task models.Task
	err := config.DB.Delete(&task, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, gin.H {
		"message" : "task deleted successfully",
	})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	err := config.DB.First(&task, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "cannot update",
		})
		return
	}
	var input models.Task
	if err_ := c.ShouldBindJSON(&input); err_ != nil {
		c.JSON(http.StatusBadRequest, err_.Error())
		return
	}
	config.DB.Model(&task).Updates(input)
	c.JSON(http.StatusOK, gin.H{
		"message" : "task updated",
	})
}