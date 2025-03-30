package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "tasks"
	CollectionUser = "users"
)

// Task represents the task entity
type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title" binding:"required"`
	Description string             `bson:"description" json:"description"`
	DueDate     time.Time          `bson:"due_date" json:"due_date"`
	Status      string             `bson:"status" json:"status"`
}

// User represents the user entity
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username" json:"username" binding:"required"`
	Password string             `bson:"password" json:"password,omitempty" binding:"required"`
	Role     string             `bson:"role" json:"role,omitempty"`
}

// TaskRepository represents the task's repository contract
type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	FetchAll(ctx context.Context) ([]Task, error)
	FetchByID(ctx context.Context, id string) (*Task, error)
	Update(ctx context.Context, id string, task *Task) error
	Delete(ctx context.Context, id string) error
}

// UserRepository represents the user's repository contract
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FetchByUsername(ctx context.Context, username string) (*User, error)
	UpdateRole(ctx context.Context, username string, role string) error
}

// TaskUsecase represents the task's usecase contract
type TaskUsecase interface {
	Create(ctx context.Context, task *Task) error
	FetchAll(ctx context.Context) ([]Task, error)
	FetchByID(ctx context.Context, id string) (*Task, error)
	Update(ctx context.Context, id string, task *Task) error
	Delete(ctx context.Context, id string) error
}

// UserUsecase represents the user's usecase contract
type UserUsecase interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, username, password string) (string, error)
	PromoteToAdmin(ctx context.Context, username string) error
}

// JWTService represents the JWT operations contract
type JWTService interface {
	GenerateToken(userID string, role string) (string, error)
	ValidateToken(token string) (map[string]interface{}, error)
}

// PasswordService represents the password operations contract
type PasswordService interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

// Response structures
type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
