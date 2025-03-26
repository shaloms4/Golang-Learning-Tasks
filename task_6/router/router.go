package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/shaloms4/Golang-Learning-Tasks/task_manager/controllers"
	"github.com/shaloms4/Golang-Learning-Tasks/task_manager/middleware"
)

// InitializeRoutes initializes the routes for the Task API.
func InitializeRoutes() {
	router := gin.Default()

	// Public routes
	router.POST("/register", controller.RegisterUser)
	router.POST("/login", controller.LoginUser)

	// Protected task routes (require authentication)
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/tasks", controller.GetTasks)
	protected.GET("/tasks/:id", controller.GetTask)

	// Admin-only routes (Create, Update, Delete tasks)
	admin := protected.Group("/")
	admin.Use(middleware.AdminMiddleware()) // Only allow admins to access

	admin.POST("/tasks", controller.AddTask)
	admin.PUT("/tasks/:id", controller.UpdateTask)
	admin.DELETE("/tasks/:id", controller.RemoveTask)

	// Admin privilege management
	admin.POST("/promote/:username", controller.PromoteUser)

	// Start the server on port 8080
	router.Run(":8080")
}
