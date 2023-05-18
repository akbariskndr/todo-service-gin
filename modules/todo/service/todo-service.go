package service

import (
	"akbariskndr/todo-service-gin/modules/todo/controller/dto"
	"akbariskndr/todo-service-gin/modules/todo/entity"
	"akbariskndr/todo-service-gin/modules/todo/repository"
)

type TodoService struct {
	Repository *repository.TodoRepository
}

func (service TodoService) FindAll(userId uint, payload *dto.FindAllTodoDto) []entity.TodoEntity {
	defaultPage := 1
	defaultLimit := 20
	if payload.Page == 0 {
		payload.Page = defaultPage
	}
	if payload.Limit == 0 {
		payload.Limit = defaultLimit
	}
	return service.Repository.FindAll(userId, payload)
}

func (service TodoService) Create(userId uint, payload *dto.CreateTodoDto) *entity.TodoEntity {
	return service.Repository.Create(userId, payload)
}

func (service TodoService) FindOne(userId uint, todoId uint) *entity.TodoEntity {
	return service.Repository.FindOne(userId, todoId)
}

func (service TodoService) Update(todoId uint, payload *dto.UpdateTodoDto, userId uint) *entity.TodoEntity {
	return service.Repository.Update(todoId, payload, userId)
}

func (service TodoService) Delete(id uint) *entity.TodoEntity {
	return service.Repository.Delete(id)
}
