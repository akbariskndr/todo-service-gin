package service

import (
	"akbariskndr/todo-service-gin/modules/todo/entity"
	"akbariskndr/todo-service-gin/modules/todo/repository"
)

type TodoService struct {
	Repository *repository.TodoRepository
}

func (service TodoService) FindAll() []entity.TodoEntity {
	return service.Repository.FindAll()
}

func (service TodoService) Create(payload entity.TodoEntity) *entity.TodoEntity {
	return service.Repository.Create(payload)
}

func (service TodoService) FindOne(id string) *entity.TodoEntity {
	return service.Repository.FindOne(id)
}

func (service TodoService) Update(id string, payload entity.TodoEntity) *entity.TodoEntity {
	return service.Repository.Update(id, payload)
}

func (service TodoService) Delete(id string) *entity.TodoEntity {
	return service.Repository.Delete(id)
}
