package usecases_test

import (
	"context"
	"testing"
	"time"

	domain "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Domain"
	"github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) FetchByUsername(ctx context.Context, username string) (*domain.User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) UpdateRole(ctx context.Context, username string, role string) error {
	args := m.Called(ctx, username, role)
	return args.Error(0)
}

// MockJWTService is a mock implementation of JWTService
type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) GenerateToken(userID string, role string) (string, error) {
	args := m.Called(userID, role)
	return args.String(0), args.Error(1)
}

func (m *MockJWTService) ValidateToken(token string) (map[string]interface{}, error) {
	args := m.Called(token)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

// MockPasswordService is a mock implementation of PasswordService
type MockPasswordService struct {
	mock.Mock
}

func (m *MockPasswordService) Hash(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordService) Compare(hashedPassword, password string) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}

type UserUsecaseSuite struct {
	suite.Suite
	mockUserRepo     *MockUserRepository
	mockJWTService   *MockJWTService
	mockPassService  *MockPasswordService
	usecase          domain.UserUsecase
	ctx              context.Context
}

func (s *UserUsecaseSuite) SetupTest() {
	s.mockUserRepo = new(MockUserRepository)
	s.mockJWTService = new(MockJWTService)
	s.mockPassService = new(MockPasswordService)
	s.usecase = usecases.NewUserUsecase(s.mockUserRepo, s.mockPassService, s.mockJWTService, 30*time.Second)
	s.ctx = context.Background()
}

func (s *UserUsecaseSuite) TestRegister() {
	user := &domain.User{
		Username: "testuser",
		Password: "password123",
		Role:     "user",
	}

	hashedPassword := "hashed_password_123"
	s.mockPassService.On("Hash", user.Password).Return(hashedPassword, nil)
	s.mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	err := s.usecase.Register(s.ctx, user)
	assert.NoError(s.T(), err)
	s.mockPassService.AssertExpectations(s.T())
	s.mockUserRepo.AssertExpectations(s.T())
}

func (s *UserUsecaseSuite) TestLogin() {
	username := "testuser"
	password := "password123"
	hashedPassword := "hashed_password_123"
	expectedToken := "jwt_token_123"

	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: hashedPassword,
		Role:     "user",
	}

	s.mockUserRepo.On("FetchByUsername", mock.Anything, username).Return(user, nil)
	s.mockPassService.On("Compare", hashedPassword, password).Return(nil)
	s.mockJWTService.On("GenerateToken", user.ID.Hex(), user.Role).Return(expectedToken, nil)

	token, err := s.usecase.Login(s.ctx, username, password)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedToken, token)
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockPassService.AssertExpectations(s.T())
	s.mockJWTService.AssertExpectations(s.T())
}

func (s *UserUsecaseSuite) TestPromoteToAdmin() {
	username := "testuser"
	s.mockUserRepo.On("UpdateRole", mock.Anything, username, "admin").Return(nil)

	err := s.usecase.PromoteToAdmin(s.ctx, username)
	assert.NoError(s.T(), err)
	s.mockUserRepo.AssertExpectations(s.T())
}

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
} 