package usecases

import (
	"context"
	"time"

	Domain "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Domain"
)

type userUsecase struct {
	userRepo        Domain.UserRepository
	passwordService Domain.PasswordService
	jwtService      Domain.JWTService
	contextTimeout  time.Duration
}

func NewUserUsecase(
	repo Domain.UserRepository,
	ps Domain.PasswordService,
	js Domain.JWTService,
	timeout time.Duration,
) Domain.UserUsecase {
	return &userUsecase{
		userRepo:        repo,
		passwordService: ps,
		jwtService:      js,
		contextTimeout:  timeout,
	}
}

func (u *userUsecase) Register(ctx context.Context, user *Domain.User) error {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	hashedPassword, err := u.passwordService.Hash(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return u.userRepo.Create(c, user)
}

func (u *userUsecase) Login(ctx context.Context, username, password string) (string, error) {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.FetchByUsername(c, username)
	if err != nil {
		return "", err
	}

	err = u.passwordService.Compare(user.Password, password)
	if err != nil {
		return "", err
	}

	return u.jwtService.GenerateToken(user.ID.Hex(), user.Role)
}

func (u *userUsecase) PromoteToAdmin(ctx context.Context, username string) error {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.userRepo.UpdateRole(c, username, "admin")
}
