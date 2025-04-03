package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Delivery/controllers"
	"github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Delivery/routers"
	infrastructure "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Infrastructure"
	Repositories "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Repositories"
	Usecases "github.com/shaloms4/Golang-Learning-Tasks/task_8/task_manager/Usecases"
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
	jwtService := infrastructure.NewJWTService()
	passwordService := infrastructure.NewPasswordService()

	// Set up usecases
	taskUsecase := Usecases.NewTaskUsecase(taskRepo, 10*time.Second)
	userUsecase := Usecases.NewUserUsecase(userRepo, passwordService, jwtService, 10*time.Second)

	// Set up controllers
	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	// Set up middleware
	authMiddleware := infrastructure.NewAuthMiddleware(jwtService)

	// Set up router
	router := routers.NewRouter(taskController, userController, authMiddleware)

	// Initialize gin engine
	engine := gin.Default()

	// Set up routes
	router.SetupRoutes(engine)

	// Start server
	log.Fatal(engine.Run(":8080"))
}
