package repositories

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	Domain "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Domain"
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
	task.ID = primitive.NewObjectID()
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

	// Create an empty map to store fields that need to be updated
	updateFields := bson.M{}

	// Dynamically add fields to the update map only if they are not empty
	if task.Title != "" {
		updateFields["title"] = task.Title
	}
	if task.Description != "" {
		updateFields["description"] = task.Description
	}
	if task.Status != "" {
		updateFields["status"] = task.Status
	}
	if task.DueDate.IsZero() { 
		updateFields["due_date"] = task.DueDate
	}

	if len(updateFields) > 0 {
		update := bson.M{
			"$set": updateFields,
		}

		// Update the task in the database
		_, err := r.db.Collection(Domain.CollectionTask).UpdateOne(ctx, bson.M{"_id": objID}, update)
		return err
	}

	// If no fields are provided to update, return nil (no update was performed)
	return nil
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
