package middleware

import (
	"net/http"
	"strings"
	"task_manager/utils"
    "fmt"
	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Missing Authorization header",
            })
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := utils.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid token: " + err.Error(),
            })
            c.Abort()
            return
        }

        // Safely extract user_id
        userIDFloat, ok := claims["user_id"].(float64)
        if !ok {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Invalid user_id type in token",
            })
            c.Abort()
            return
        }

        userID := uint(userIDFloat)
        fmt.Println("User ID from token:", userID)

        c.Set("user_id", userID)
        c.Next()
    }
}
