package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/shaloms4/Golang-Learning-Tasks/task_manager/controllers"
)

// InitializeRoutes initializes the routes for the Task API.
func InitializeRoutes() {
	router := gin.Default()

	// Define routes for CRUD operations on tasks
	router.GET("/tasks", controller.GetTasks)
	router.GET("/tasks/:id", controller.GetTask)
	router.POST("/tasks", controller.AddTask)
	router.PUT("/tasks/:id", controller.UpdateTask)
	router.DELETE("/tasks/:id", controller.RemoveTask)

	// Start the server on port 8080
	router.Run(":8080")
}
