package module

import (
	"akbariskndr/todo-service-gin/database"
	"akbariskndr/todo-service-gin/modules/auth/controller"
	"akbariskndr/todo-service-gin/modules/auth/repository"
	"akbariskndr/todo-service-gin/modules/auth/service"
)

type AuthModule struct {
	Repository  *repository.UserRepository
	UserService *service.AuthService
	Controller  *controller.AuthController
}

func InitModule() *AuthModule {
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
