package repositories

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"

	Domain "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Domain"
)

type userRepository struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewUserRepository() (Domain.UserRepository, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGODB_URL")

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "task_manager"
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	return &userRepository{
		client: client,
		db:     db,
	}, nil
}

func (r *userRepository) Create(ctx context.Context, user *Domain.User) error {
	// Check if this is the first user
	count, err := r.db.Collection(Domain.CollectionUser).CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	// If this is the first user, make them an admin
	if count == 0 {
		user.Role = "admin"
	} else if user.Role == "" {
		user.Role = "user" // Default role for non-first users
	}

	user.ID = primitive.NewObjectID()
	_, err = r.db.Collection(Domain.CollectionUser).InsertOne(ctx, user)
	return err
}

func (r *userRepository) FetchByUsername(ctx context.Context, username string) (*Domain.User, error) {
	var user Domain.User
	err := r.db.Collection(Domain.CollectionUser).FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateRole(ctx context.Context, username string, role string) error {
	update := bson.M{
		"$set": bson.M{
			"role": role,
		},
	}

	result, err := r.db.Collection(Domain.CollectionUser).UpdateOne(ctx, bson.M{"username": username}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
