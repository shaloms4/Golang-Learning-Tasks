package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Delivery/controllers"
	infrastructure "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Infrastructure"
)

type Router struct {
	taskController *controllers.TaskController
	userController *controllers.UserController
	authMiddleware *infrastructure.AuthMiddleware
}

func NewRouter(
	taskController *controllers.TaskController,
	userController *controllers.UserController,
	authMiddleware *infrastructure.AuthMiddleware,
) *Router {
	return &Router{
		taskController: taskController,
		userController: userController,
		authMiddleware: authMiddleware,
	}
}

func (r *Router) SetupRoutes(engine *gin.Engine) {
	// Public routes
	engine.POST("/register", r.userController.Register)
	engine.POST("/login", r.userController.Login)

	// Protected routes
	protected := engine.Group("")
	protected.Use(r.authMiddleware.Middleware())

	// Task routes
	protected.POST("/tasks", infrastructure.RequireAdmin(), r.taskController.CreateTask)
	protected.GET("/tasks", r.taskController.FetchAllTasks)
	protected.GET("/tasks/:id", r.taskController.FetchTaskByID)
	protected.PUT("/tasks/:id", infrastructure.RequireAdmin(), r.taskController.UpdateTask)
	protected.DELETE("/tasks/:id", infrastructure.RequireAdmin(), r.taskController.DeleteTask)

	// Admin routes
	protected.PUT("/users/:username/promote", infrastructure.RequireAdmin(), r.userController.PromoteToAdmin)
}