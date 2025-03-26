package data

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shaloms4/Golang-Learning-Tasks/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global variable to store the collection
var taskCollection *mongo.Collection

// InitializeMongoDB sets up the MongoDB connection
func InitializeMongoDB() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	clientOptions := options.Client().ApplyURI(uri)

	// Establish the MongoDB connection
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	// Initialize both task and user collections
	db := client.Database("taskmanager")
	taskCollection = db.Collection("tasks")
	userCollection = db.Collection("users") 

	log.Println("Connected to MongoDB!")
}

// GetTasks retrieves all tasks from MongoDB
func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := taskCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// GetTaskByID retrieves a single task by ID
func GetTaskByID(id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID format")
	}

	var task models.Task
	err = taskCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	return &task, nil
}

// AddTask inserts a new task into MongoDB
func AddTask(newTask models.Task) (*mongo.InsertOneResult, error) {
	// Insert task into MongoDB collection
	result, err := taskCollection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTask updates an existing task by its ID, but only updates the fields provided in the request.
func UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	// Convert string ID to MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID format")
	}

	// Prepare the update operation
	update := bson.M{"$set": bson.M{}}

	// Conditionally add fields to the update object if they are not empty
	if updatedTask.Title != "" {
		update["$set"].(bson.M)["title"] = updatedTask.Title
	}
	if updatedTask.Description != "" {
		update["$set"].(bson.M)["description"] = updatedTask.Description
	}
	if updatedTask.Status != "" {
		update["$set"].(bson.M)["status"] = updatedTask.Status
	}
	if !updatedTask.DueDate.IsZero() {
		update["$set"].(bson.M)["due_date"] = updatedTask.DueDate
	}

	// Execute the update operation
	result := taskCollection.FindOneAndUpdate(context.TODO(), bson.M{"_id": objID}, update)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		}
		return nil, result.Err()
	}

	// Retrieve the updated task from the database
	return GetTaskByID(id)
}

// RemoveTask deletes a task
func RemoveTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID format")
	}

	result, err := taskCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
