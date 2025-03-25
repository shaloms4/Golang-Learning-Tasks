package main

import (
	"github.com/shaloms4/Golang-Learning-Tasks/task_manager/data"
	"github.com/shaloms4/Golang-Learning-Tasks/task_manager/router"
)

func main() {
	// Initialize MongoDB connection
	data.InitializeMongoDB()
	// Initialize routes and start the server
	router.InitializeRoutes()
}
