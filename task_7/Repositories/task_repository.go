package repository

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	Domain "github.com/shaloms4/Golang-Learning-Tasks/task_7/Domain"
)

type taskRepository struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewTaskRepository() (Domain.TaskRepository, error) {
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
	return &taskRepository{
		client: client,
		db:     db,
	}, nil
}

func (r *taskRepository) Create(ctx context.Context, task *Domain.Task) error {
	_, err := r.db.Collection(Domain.CollectionTask).InsertOne(ctx, task)
	return err
}

func (r *taskRepository) FetchAll(ctx context.Context) ([]Domain.Task, error) {
	var tasks []Domain.Task
	cursor, err := r.db.Collection(Domain.CollectionTask).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) FetchByID(ctx context.Context, id string) (*Domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var task Domain.Task
	err = r.db.Collection(Domain.CollectionTask).FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepository) Update(ctx context.Context, id string, task *Domain.Task) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"status":      task.Status,
			"due_date":    task.DueDate,
		},
	}

	_, err = r.db.Collection(Domain.CollectionTask).UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func (r *taskRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := r.db.Collection(Domain.CollectionTask).DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
