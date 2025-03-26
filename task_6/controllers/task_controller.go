package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaloms4/Golang-Learning-Tasks/task_manager/data"
	"github.com/shaloms4/Golang-Learning-Tasks/task_manager/models"
)

// GetTasks retrieves all tasks
func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTask retrieves a task by ID
func GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// AddTask adds a new task
func AddTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate that title is provided
	if newTask.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	result, err := data.AddTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created", "task_id": result.InsertedID})
}

// UpdateTask updates an existing task by its ID, but only updates the fields provided in the request.
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task

	// Bind the request JSON to the updatedTask object
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to update the task with the new data
	task, err := data.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Return the updated task in response
	c.JSON(http.StatusOK, gin.H{"message": "Task updated", "task": task})
}

// RemoveTask removes a task by ID
func RemoveTask(c *gin.Context) {
	id := c.Param("id")
	err := data.RemoveTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}

// RegisterUser registers a new user
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := data.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// LoginUser logs in a user
func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := data.LoginUser(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// PromoteUser promotes a regular user to an admin
func PromoteUser(c *gin.Context) {
	username := c.Param("username")

	err := data.PromoteToAdmin(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}
