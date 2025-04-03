package repositories_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	domain "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Domain"
	"github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Repositories"
)

type TaskRepositorySuite struct {
	suite.Suite
	repository domain.TaskRepository
	client     *mongo.Client
	db         *mongo.Database
	ctx        context.Context
}

func (s *TaskRepositorySuite) SetupSuite() {
	os.Setenv("MONGODB_URL", "mongodb://localhost:27017")
	os.Setenv("DB_NAME", "task_manager_test")

	s.ctx = context.Background()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	client, err := mongo.Connect(s.ctx, clientOptions)
	assert.NoError(s.T(), err)
	s.client = client

	s.db = client.Database(os.Getenv("DB_NAME"))

	repo, err := repositories.NewTaskRepository()
	assert.NoError(s.T(), err)
	s.repository = repo
}

func (s *TaskRepositorySuite) TearDownSuite() {
	err := s.db.Drop(s.ctx)
	assert.NoError(s.T(), err)

	err = s.client.Disconnect(s.ctx)
	assert.NoError(s.T(), err)
}

func (s *TaskRepositorySuite) TestCreateTask() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}

	err := s.repository.Create(s.ctx, task)
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), task.ID)
}

func (s *TaskRepositorySuite) TestFetchAllTasks() {
	task1 := &domain.Task{
		Title:       "Task 1",
		Description: "Description 1",
		Status:      "pending",
	}
	task2 := &domain.Task{
		Title:       "Task 2",
		Description: "Description 2",
		Status:      "completed",
	}

	err := s.repository.Create(s.ctx, task1)
	assert.NoError(s.T(), err)
	err = s.repository.Create(s.ctx, task2)
	assert.NoError(s.T(), err)

	tasks, err := s.repository.FetchAll(s.ctx)
	assert.NoError(s.T(), err)
	assert.GreaterOrEqual(s.T(), len(tasks), 2)
}

func (s *TaskRepositorySuite) TestFetchTaskByID() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}

	err := s.repository.Create(s.ctx, task)
	assert.NoError(s.T(), err)

	fetchedTask, err := s.repository.FetchByID(s.ctx, task.ID.Hex())
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), task.Title, fetchedTask.Title)
	assert.Equal(s.T(), task.Description, fetchedTask.Description)
}

func (s *TaskRepositorySuite) TestUpdateTask() {
	task := &domain.Task{
		Title:       "Original Title",
		Description: "Original Description",
		Status:      "pending",
	}

	err := s.repository.Create(s.ctx, task)
	assert.NoError(s.T(), err)

	updatedTask := &domain.Task{
		Title:       "Updated Title",
		Description: "Updated Description",
		Status:      "completed",
	}

	err = s.repository.Update(s.ctx, task.ID.Hex(), updatedTask)
	assert.NoError(s.T(), err)

	// Fetch and verify the update
	fetchedTask, err := s.repository.FetchByID(s.ctx, task.ID.Hex())
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), updatedTask.Title, fetchedTask.Title)
	assert.Equal(s.T(), updatedTask.Description, fetchedTask.Description)
	assert.Equal(s.T(), updatedTask.Status, fetchedTask.Status)
}

func (s *TaskRepositorySuite) TestDeleteTask() {
	// Create a test task
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}

	err := s.repository.Create(s.ctx, task)
	assert.NoError(s.T(), err)

	// Delete the task
	err = s.repository.Delete(s.ctx, task.ID.Hex())
	assert.NoError(s.T(), err)

	_, err = s.repository.FetchByID(s.ctx, task.ID.Hex())
	assert.Error(s.T(), err)
}

func TestTaskRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskRepositorySuite))
} 