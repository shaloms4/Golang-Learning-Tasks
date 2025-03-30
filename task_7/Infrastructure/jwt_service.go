package Infrastructure

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	Domain "github.com/shaloms4/Golang-Learning-Tasks/task_7/Domain"
)

type jwtService struct {
	secretKey string
}

func NewJWTService() Domain.JWTService {
	return &jwtService{
		secretKey: os.Getenv("JWT_SECRET"),
	}
}

func (s *jwtService) GenerateToken(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) ValidateToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}
