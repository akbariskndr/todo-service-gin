package module

import (
	"akbariskndr/todo-service-gin/database"
	"akbariskndr/todo-service-gin/modules/auth/controller"
	"akbariskndr/todo-service-gin/modules/auth/repository"
	"akbariskndr/todo-service-gin/modules/auth/service"
	"sync"
)

type AuthModule struct {
	Repository  *repository.UserRepository
	AuthService *service.AuthService
	UserService *service.UserService
	Controller  *controller.AuthController
}

var lock = &sync.Mutex{}

var singleton *AuthModule

func createInstance() *AuthModule {
	var userRepository *repository.UserRepository = &repository.UserRepository{
		DB: database.Connector,
	}
	var authService *service.AuthService = &service.AuthService{
		Repository: userRepository,
	}
	var userService *service.UserService = &service.UserService{
		Repository: userRepository,
	}
	var authController *controller.AuthController = &controller.AuthController{
		AuthService: authService,
		UserService: userService,
	}

	return &AuthModule{
		Repository:  userRepository,
		UserService: userService,
		AuthService: authService,
		Controller:  authController,
	}
}

func GetInstance() *AuthModule {
	if singleton == nil {
		lock.Lock()

		defer lock.Unlock()

		if singleton == nil {
			singleton = createInstance()
		}
	}

	return singleton
}
