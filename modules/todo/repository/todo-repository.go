package repository

import (
	"akbariskndr/todo-service-gin/modules/todo/entity"

	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func (repository TodoRepository) FindAll() []entity.TodoEntity {
	var todos []entity.TodoEntity
	repository.DB.Find(&todos)

	return todos
}

func (repository TodoRepository) Create(payload entity.TodoEntity) *entity.TodoEntity {
	var todo *entity.TodoEntity

	repository.DB.Create(todo)

	return todo
}

func (repository TodoRepository) FindOne(id string) *entity.TodoEntity {
	var todo *entity.TodoEntity

	if err := repository.DB.First(&todo, id).Error; err != nil {
		return nil
	}

	return todo
}

func (repository TodoRepository) Update(id string, payload entity.TodoEntity) *entity.TodoEntity {
	var todo *entity.TodoEntity

	if err := repository.DB.First(&todo, id).Error; err != nil {
		return nil
	}

	todo.Title = payload.Title
	todo.Completed = payload.Completed
	repository.DB.Save(todo)

	return todo
}

func (repository TodoRepository) Delete(id string) *entity.TodoEntity {
	var todo *entity.TodoEntity

	if err := repository.DB.First(&todo, id).Error; err != nil {
		return nil
	}
	repository.DB.Delete(todo)

	return todo
}
