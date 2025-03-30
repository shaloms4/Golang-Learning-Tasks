package main

import (
	"log"
	"time"

	"github.com/shaloms4/Golang-Learning-Tasks/task_7/Delivery/controllers"
	"github.com/shaloms4/Golang-Learning-Tasks/task_7/Delivery/routers"
	"github.com/shaloms4/Golang-Learning-Tasks/task_7/Infrastructure"
	Repositories "github.com/shaloms4/Golang-Learning-Tasks/task_7/Repositories"
	Usecases "github.com/shaloms4/Golang-Learning-Tasks/task_7/Usecases"
)

func main() {
	// Set up repositories
	taskRepo, err := Repositories.NewTaskRepository()
	if err != nil {
		log.Fatal(err)
	}

	userRepo, err := Repositories.NewUserRepository()
	if err != nil {
		log.Fatal(err)
	}

	// Set up services
	jwtService := Infrastructure.NewJWTService()
	passwordService := Infrastructure.NewPasswordService()

	// Set up usecases
	taskUsecase := Usecases.NewTaskUsecase(taskRepo, 10*time.Second)
	userUsecase := Usecases.NewUserUsecase(userRepo, passwordService, jwtService, 10*time.Second)

	// Set up controllers
	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	// Set up middleware
	authMiddleware := Infrastructure.NewAuthMiddleware(jwtService)

	// Set up router
	r := routers.NewRouter(taskController, userController, authMiddleware)

	// Start server
	log.Fatal(r.Run(":8080"))
}
