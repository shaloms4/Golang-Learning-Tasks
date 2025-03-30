package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/shaloms4/Golang-Learning-Tasks/task_7/Delivery/controllers"
	"github.com/shaloms4/Golang-Learning-Tasks/task_7/Infrastructure"
)

func NewRouter(
	taskController *controllers.TaskController,
	userController *controllers.UserController,
	authMiddleware *Infrastructure.AuthMiddleware,
) *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	// Protected routes
	protected := router.Group("")
	protected.Use(authMiddleware.Middleware())
	{
		// Task routes that all authenticated users can access
		protected.GET("/tasks", taskController.FetchAllTasks)
		protected.GET("/tasks/:id", taskController.FetchTaskByID)

		// Admin-only task routes
		admin := protected.Group("")
		admin.Use(Infrastructure.RequireAdmin())
		{
			admin.POST("/tasks", taskController.CreateTask)
			admin.PUT("/tasks/:id", taskController.UpdateTask)
			admin.DELETE("/tasks/:id", taskController.DeleteTask)
			admin.PUT("/users/:username/promote", userController.PromoteToAdmin)
		}
	}

	return router
}