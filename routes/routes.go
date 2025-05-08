package routes

import (
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Task Manager API",
		})
	})
	//Auth routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/dashboard", func(c *gin.Context) {
			user, _ := c.Get("user")
			c.JSON(200, gin.H{
				"message": "Welcome to dashboard!",
				"user":    user,
			})
		})
	}
}
