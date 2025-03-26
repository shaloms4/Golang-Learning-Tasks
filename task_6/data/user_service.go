package data

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/shaloms4/Golang-Learning-Tasks/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

// Hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Verify password
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Register user
func RegisterUser(user models.User) (*models.User, error) {
	var existingUser models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return nil, errors.New("username already taken")
	}

	user.ID = primitive.NewObjectID()
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	count, _ := userCollection.CountDocuments(context.TODO(), bson.M{})
	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	_, err = userCollection.InsertOne(context.TODO(), user)
	return &user, err
}

// Login user
func LoginUser(username, password string) (string, error) {
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := CheckPassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// PromoteToAdmin upgrades a user to an admin role
func PromoteToAdmin(username string) error {
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	result, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}
