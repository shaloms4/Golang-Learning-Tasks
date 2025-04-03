package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"

	Domain "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Domain"
	usecases "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Usecases"
)

// MockTaskRepository is a mock implementation of TaskRepository
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(ctx context.Context, task *Domain.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func (m *MockTaskRepository) FetchAll(ctx context.Context) ([]Domain.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) FetchByID(ctx context.Context, id string) (*Domain.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) Update(ctx context.Context, id string, task *Domain.Task) error {
	args := m.Called(ctx, id, task)
	return args.Error(0)
}

func (m *MockTaskRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type TaskUsecaseSuite struct {
	suite.Suite
	mockRepo *MockTaskRepository
	usecase  Domain.TaskUsecase
	ctx      context.Context
}

func (s *TaskUsecaseSuite) SetupTest() {
	s.mockRepo = new(MockTaskRepository)
	s.usecase = usecases.NewTaskUsecase(s.mockRepo, 30*time.Second)
	s.ctx = context.Background()
}

func (s *TaskUsecaseSuite) TestCreateTask() {
	task := &Domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}

	s.mockRepo.On("Create", mock.Anything, task).Return(nil)

	err := s.usecase.Create(s.ctx, task)
	assert.NoError(s.T(), err)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUsecaseSuite) TestFetchAllTasks() {
	expectedTasks := []Domain.Task{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 1",
			Description: "Description 1",
			Status:      "pending",
		},
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 2",
			Description: "Description 2",
			Status:      "completed",
		},
	}

	s.mockRepo.On("FetchAll", mock.Anything).Return(expectedTasks, nil)

	tasks, err := s.usecase.FetchAll(s.ctx)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedTasks, tasks)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUsecaseSuite) TestFetchTaskByID() {
	taskID := primitive.NewObjectID().Hex()
	expectedTask := &Domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "pending",
	}

	s.mockRepo.On("FetchByID", mock.Anything, taskID).Return(expectedTask, nil)

	task, err := s.usecase.FetchByID(s.ctx, taskID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedTask, task)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUsecaseSuite) TestUpdateTask() {
	taskID := primitive.NewObjectID().Hex()
	task := &Domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "completed",
	}

	s.mockRepo.On("Update", mock.Anything, taskID, task).Return(nil)

	err := s.usecase.Update(s.ctx, taskID, task)
	assert.NoError(s.T(), err)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUsecaseSuite) TestDeleteTask() {
	taskID := primitive.NewObjectID().Hex()

	s.mockRepo.On("Delete", mock.Anything, taskID).Return(nil)

	err := s.usecase.Delete(s.ctx, taskID)
	assert.NoError(s.T(), err)
	s.mockRepo.AssertExpectations(s.T())
}

func TestTaskUsecaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseSuite))
} 