package module

import (
	"akbariskndr/todo-service-gin/database"
	"akbariskndr/todo-service-gin/modules/todo/controller"
	"akbariskndr/todo-service-gin/modules/todo/repository"
	"akbariskndr/todo-service-gin/modules/todo/service"
	"sync"
)

type TodoModule struct {
	Repository *repository.TodoRepository
	Service    *service.TodoService
	Controller *controller.TodoController
}

var lock = &sync.Mutex{}

var singleton *TodoModule

func createInstance() *TodoModule {
	var todoRepository *repository.TodoRepository = &repository.TodoRepository{
		DB: database.Connector,
	}
	var todoService *service.TodoService = &service.TodoService{
		Repository: todoRepository,
	}
	var todoController *controller.TodoController = &controller.TodoController{
		Service: todoService,
	}
	return &TodoModule{
		Repository: todoRepository,
		Service:    todoService,
		Controller: todoController,
	}
}

func GetInstance() *TodoModule {
	if singleton == nil {
		lock.Lock()

		defer lock.Unlock()

		if singleton == nil {
			singleton = createInstance()
		}
	}

	return singleton
}
