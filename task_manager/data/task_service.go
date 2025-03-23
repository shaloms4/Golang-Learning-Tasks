package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/shaloms4/Golang-Learning-Tasks/task_manager/models"
)

var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

// GetTasks retrieves all tasks from the in-memory database.
func GetTasks() []models.Task {
	return tasks
}

// GetTaskByID retrieves a single task by its ID.
func GetTaskByID(id string) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

// AddTask adds a new task to the in-memory database.
func AddTask(newTask models.Task) {
	tasks = append(tasks, newTask)
}

// UpdateTask updates an existing task by its ID.
func UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			// Only update fields that were provided
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			if !updatedTask.DueDate.IsZero() {
				tasks[i].DueDate = updatedTask.DueDate
			}

			return &tasks[i], nil
		}
	}
	return nil, fmt.Errorf("task with ID %s not found", id)
}

// RemoveTask deletes a task by its ID.
func RemoveTask(id string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
