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

type UserRepositorySuite struct {
	suite.Suite
	repository domain.UserRepository
	client     *mongo.Client
	db         *mongo.Database
	ctx        context.Context
}

func (s *UserRepositorySuite) SetupSuite() {
	os.Setenv("MONGODB_URL", "mongodb://localhost:27017")
	os.Setenv("DB_NAME", "task_manager_test")

	s.ctx = context.Background()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	client, err := mongo.Connect(s.ctx, clientOptions)
	assert.NoError(s.T(), err)
	s.client = client

	s.db = client.Database(os.Getenv("DB_NAME"))

	// Create repository
	repo, err := repositories.NewUserRepository()
	assert.NoError(s.T(), err)
	s.repository = repo
}

func (s *UserRepositorySuite) TearDownSuite() {
	err := s.db.Drop(s.ctx)
	assert.NoError(s.T(), err)

	err = s.client.Disconnect(s.ctx)
	assert.NoError(s.T(), err)
}

func (s *UserRepositorySuite) TestCreateUser() {
	user := &domain.User{
		Username: "testuser",
		Password: "password123",
		Role:     "user",
	}

	err := s.repository.Create(s.ctx, user)
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), user.ID)
}

func (s *UserRepositorySuite) TestFetchByUsername() {
	// First create a user
	user := &domain.User{
		Username: "testuser_fetch",
		Password: "password123",
		Role:     "user",
	}
	err := s.repository.Create(s.ctx, user)
	assert.NoError(s.T(), err)

	// Then fetch it by username
	fetchedUser, err := s.repository.FetchByUsername(s.ctx, user.Username)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), user.Username, fetchedUser.Username)
	assert.Equal(s.T(), user.Role, fetchedUser.Role)
}

func (s *UserRepositorySuite) TestUpdateRole() {
	user := &domain.User{
		Username: "testuser_role",
		Password: "password123",
		Role:     "user",
	}
	err := s.repository.Create(s.ctx, user)
	assert.NoError(s.T(), err)

	// Update the role
	err = s.repository.UpdateRole(s.ctx, user.Username, "admin")
	assert.NoError(s.T(), err)

	// Verify the role update
	updatedUser, err := s.repository.FetchByUsername(s.ctx, user.Username)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "admin", updatedUser.Role)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
} 