package controllers

import (
	"net/http"
	"fmt"
	"task_manager/config"
	"task_manager/models"
	"task_manager/utils"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c* gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error ; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func RegisterUser(c *gin.Context) {
    var Newuser models.User
    if err := c.ShouldBindJSON(&Newuser); err != nil {
        c.JSON(http.StatusBadRequest, err.Error())
        return
    }

    // Hash password
    hashedPassword, err := utils.HashPassword(Newuser.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    Newuser.Password = hashedPassword

    // Save user with hashed password
    if err := config.DB.Create(&Newuser).Error; err != nil {
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "user registered",
    })
}

func LoginUser(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind login input
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user models.User
	if err := config.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check password (plain-text comparison for now)
	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
    return
	}

	token, err := utils.GenerateTokens(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "couldn't generate token",
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token" : token,
	})
}

func GetUserProfile(c* gin.Context) {
	userIdRaw, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "user id not found",
		})
		return
	}
	userId, er := userIdRaw.(uint)
	fmt.Println("UserID from context:", userId)

	if !er {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "user id of invalid type",
		})
		return
	}
	var user models.User
	err := config.DB.First(&user, userId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user" : user,
	})
}

func UpdateUser(c* gin.Context) {
	UserId := c.MustGet("user_id").(uint)
	var user models.User
	err := config.DB.First(&user, UserId).Error ;
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "cannot find User",
		})
		return
	}
	var input struct {
		Username string
		Email string
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	user.Username = input.Username
	user.Email = input.Email
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "cannot update user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "user updated",
	})
}

func DeleteUser(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	if err := config.DB.Delete(&models.User{}, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
