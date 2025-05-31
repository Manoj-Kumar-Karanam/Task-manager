package routers

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(u* gin.Engine) {
	public := u.Group("/auth")
	{
		public.POST("/register", controllers.RegisterUser)
		public.POST("/login", controllers.LoginUser)
	} 
	protected := u.Group("/user") 
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", controllers.GetUserProfile)
		protected.PUT("/update", controllers.UpdateUser)
		protected.DELETE("/delete", controllers.DeleteUser)
	}
}