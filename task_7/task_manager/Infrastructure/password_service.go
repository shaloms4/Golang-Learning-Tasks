package Infrastructure

import (
	"golang.org/x/crypto/bcrypt"

	Domain "github.com/shaloms4/Golang-Learning-Tasks/task_7/Domain"
)

type PasswordService struct{}

func NewPasswordService() Domain.PasswordService {
	return &PasswordService{}
}

func (s *PasswordService) Hash(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (s *PasswordService) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
