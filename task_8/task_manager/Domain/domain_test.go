package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	Domain "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Domain"
)

func TestTaskModel(t *testing.T) {
	t.Run("Task Creation", func(t *testing.T) {
		task := &Domain.Task{
			Title:       "Test Task",
			Description: "Test Description",
			DueDate:     time.Now().Add(24 * time.Hour),
			Status:      "pending",
		}

		assert.NotEmpty(t, task.Title, "Title should not be empty")
		assert.Equal(t, "Test Task", task.Title, "Title should match")
		assert.Equal(t, "Test Description", task.Description, "Description should match")
		assert.Equal(t, "pending", task.Status, "Status should be pending")
	})

	t.Run("Task Validation", func(t *testing.T) {
		task := &Domain.Task{
			Title: "", 
		}

		assert.Empty(t, task.Title, "Title should be empty")
	})
}

func TestUserModel(t *testing.T) {
	t.Run("User Creation", func(t *testing.T) {
		user := &Domain.User{
			Username: "testuser",
			Password: "password123",
			Role:     "user",
		}

		assert.NotEmpty(t, user.Username, "Username should not be empty")
		assert.Equal(t, "testuser", user.Username, "Username should match")
		assert.Equal(t, "password123", user.Password, "Password should match")
		assert.Equal(t, "user", user.Role, "Role should be user")
	})

	t.Run("User Validation", func(t *testing.T) {
		user := &Domain.User{
			Username: "", 
			Password: "", 
		}

		assert.Empty(t, user.Username, "Username should be empty")
		assert.Empty(t, user.Password, "Password should be empty")
	})
}

func TestResponseStructures(t *testing.T) {
	t.Run("ErrorResponse", func(t *testing.T) {
		errorResp := Domain.ErrorResponse{
			Message: "Test error message",
		}

		assert.Equal(t, "Test error message", errorResp.Message, "Error message should match")
	})

	t.Run("SuccessResponse", func(t *testing.T) {
		successResp := Domain.SuccessResponse{
			Message: "Operation successful",
		}

		assert.Equal(t, "Operation successful", successResp.Message, "Success message should match")
	})

	t.Run("LoginResponse", func(t *testing.T) {
		loginResp := Domain.LoginResponse{
			Token: "test-token-123",
		}

		assert.Equal(t, "test-token-123", loginResp.Token, "Token should match")
	})
} 