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
	UserService *service.AuthService
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
	var authController *controller.AuthController = &controller.AuthController{
		Service: authService,
	}

	return &AuthModule{
		Repository:  userRepository,
		UserService: authService,
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
