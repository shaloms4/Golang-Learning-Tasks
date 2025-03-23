package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/shaloms4/Golang-Learning-Tasks/task_manager/data"
	"github.com/shaloms4/Golang-Learning-Tasks/task_manager/models"
)

// GetTasks retrieves all tasks
func GetTasks(c *gin.Context) {
	tasks := service.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

// GetTask retrieves a task by ID
func GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := service.GetTaskByID(id)
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

	service.AddTask(newTask)
	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

// UpdateTask updates an existing task
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := service.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated", "task": task})
}

// RemoveTask removes a task by ID
func RemoveTask(c *gin.Context) {
	id := c.Param("id")
	err := service.RemoveTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task removed"})
}
