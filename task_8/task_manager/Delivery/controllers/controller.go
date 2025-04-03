package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	Domain "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Domain"
)

type TaskController struct {
	taskUsecase Domain.TaskUsecase
}

type UserController struct {
	userUsecase Domain.UserUsecase
}

func NewTaskController(usecase Domain.TaskUsecase) *TaskController {
	return &TaskController{
		taskUsecase: usecase,
	}
}

func NewUserController(usecase Domain.UserUsecase) *UserController {
	return &UserController{
		userUsecase: usecase,
	}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	// Check if user is admin
	role, exists := ctx.Get("userRole")
	if !exists || role != "admin" {
		ctx.JSON(http.StatusForbidden, Domain.ErrorResponse{Message: "Only admins can create tasks"})
		return
	}

	var task Domain.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, Domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := c.taskUsecase.Create(ctx.Request.Context(), &task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Domain.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, Domain.SuccessResponse{Message: "Task created successfully"})
}

// FetchAllTasks - All authenticated users can fetch tasks
func (c *TaskController) FetchAllTasks(ctx *gin.Context) {
	tasks, err := c.taskUsecase.FetchAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Domain.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) FetchTaskByID(ctx *gin.Context) {
	taskID := ctx.Param("id")
	task, err := c.taskUsecase.FetchByID(ctx.Request.Context(), taskID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, Domain.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	// Check if user is admin
	role, exists := ctx.Get("userRole")
	if !exists || role != "admin" {
		ctx.JSON(http.StatusForbidden, Domain.ErrorResponse{Message: "Only admins can update tasks"})
		return
	}

	taskID := ctx.Param("id")
	var task Domain.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, Domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := c.taskUsecase.Update(ctx.Request.Context(), taskID, &task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Domain.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Domain.SuccessResponse{Message: "Task updated successfully"})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	// Check if user is admin
	role, exists := ctx.Get("userRole")
	if !exists || role != "admin" {
		ctx.JSON(http.StatusForbidden, Domain.ErrorResponse{Message: "Only admins can delete tasks"})
		return
	}

	taskID := ctx.Param("id")
	err := c.taskUsecase.Delete(ctx.Request.Context(), taskID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Domain.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Domain.SuccessResponse{Message: "Task deleted successfully"})
}

func (c *UserController) Register(ctx *gin.Context) {
	var user Domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, Domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := c.userUsecase.Register(ctx.Request.Context(), &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Domain.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, Domain.SuccessResponse{Message: "User registered successfully"})
}

func (c *UserController) Login(ctx *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, Domain.ErrorResponse{Message: err.Error()})
		return
	}

	token, err := c.userUsecase.Login(ctx.Request.Context(), credentials.Username, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, Domain.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Domain.LoginResponse{Token: token})
}

func (c *UserController) PromoteToAdmin(ctx *gin.Context) {
	username := ctx.Param("username")
	err := c.userUsecase.PromoteToAdmin(ctx.Request.Context(), username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Domain.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Domain.SuccessResponse{Message: "User promoted to admin successfully"})
}