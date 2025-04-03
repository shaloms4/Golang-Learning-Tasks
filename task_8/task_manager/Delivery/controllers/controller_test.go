package controllers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	Domain "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Domain"
	"github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Delivery/controllers"
)

type MockTaskUsecase struct {
	mock.Mock
}

func (m *MockTaskUsecase) Create(ctx context.Context, task *Domain.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func (m *MockTaskUsecase) FetchAll(ctx context.Context) ([]Domain.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]Domain.Task), args.Error(1)
}

func (m *MockTaskUsecase) FetchByID(ctx context.Context, id string) (*Domain.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskUsecase) Update(ctx context.Context, id string, task *Domain.Task) error {
	args := m.Called(ctx, id, task)
	return args.Error(0)
}

func (m *MockTaskUsecase) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type ControllerSuite struct {
	suite.Suite
	mockTaskUsecase *MockTaskUsecase
	taskController  *controllers.TaskController
	router          *gin.Engine
}

func (s *ControllerSuite) SetupTest() {
	s.mockTaskUsecase = new(MockTaskUsecase)
	s.taskController = controllers.NewTaskController(s.mockTaskUsecase)
	
	gin.SetMode(gin.TestMode)
	s.router = gin.New()

	s.router.Use(func(c *gin.Context) {
		c.Set("userRole", "admin")
		c.Next()
	})

	s.router.POST("/tasks", s.taskController.CreateTask)
	s.router.GET("/tasks", s.taskController.FetchAllTasks)
	s.router.GET("/tasks/:id", s.taskController.FetchTaskByID)
	s.router.PUT("/tasks/:id", s.taskController.UpdateTask)
	s.router.DELETE("/tasks/:id", s.taskController.DeleteTask)
}

func (s *ControllerSuite) TestCreateTask() {
	task := &Domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}

	s.mockTaskUsecase.On("Create", mock.Anything, task).Return(nil)

	w := httptest.NewRecorder()
	jsonBody := `{"title":"Test Task","description":"Test Description","status":"pending"}`
	req, _ := http.NewRequest("POST", "/tasks", strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusCreated, w.Code)
	s.mockTaskUsecase.AssertExpectations(s.T())
}

func (s *ControllerSuite) TestFetchAllTasks() {
	tasks := []Domain.Task{
		{Title: "Task 1", Description: "Description 1", Status: "pending"},
		{Title: "Task 2", Description: "Description 2", Status: "completed"},
	}

	s.mockTaskUsecase.On("FetchAll", mock.Anything).Return(tasks, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	s.mockTaskUsecase.AssertExpectations(s.T())
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerSuite))
} 