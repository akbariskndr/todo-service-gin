package repository

import (
	"akbariskndr/todo-service-gin/modules/todo/controller/dto"
	"akbariskndr/todo-service-gin/modules/todo/entity"
	"fmt"

	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func (repository TodoRepository) FindAll(userId uint, payload *dto.FindAllTodoDto) []entity.TodoEntity {
	var todos []entity.TodoEntity

	query := repository.DB.Where("user_id = ?", userId)

	if payload.Search != "" {
		query = query.Where("title LIKE ?", fmt.Sprintf("%%%s%%", payload.Search))
	}

	if payload.Sort != "" && payload.Order != "" {
		query = query.Order(fmt.Sprintf("%s %s", payload.Sort, payload.Order))
	} else {
		query = query.Order("created_at DESC")
	}

	query.Offset((payload.Page * payload.Limit) - payload.Limit).Limit(payload.Limit).Find(&todos)

	if todos == nil {
		return []entity.TodoEntity{}
	}

	return todos
}

func (repository TodoRepository) Create(id uint, payload *dto.CreateTodoDto) *entity.TodoEntity {
	todo := &entity.TodoEntity{
		Title:     payload.Title,
		Completed: false,
		UserId:    id,
	}

	repository.DB.Create(todo)

	return todo
}

func (repository TodoRepository) FindOne(userId uint, todoId uint) *entity.TodoEntity {
	var todo *entity.TodoEntity

	res := repository.DB.Where("id = ?", todoId).Where("user_id = ?", userId).First(&todo, userId)
	if err := res.Error; err != nil {
		return nil
	}

	return todo
}

func (repository TodoRepository) Update(todoId uint, payload *dto.UpdateTodoDto, userId uint) *entity.TodoEntity {
	var todo *entity.TodoEntity

	res := repository.DB.Where("id = ?", todoId).Where("user_id = ?", userId).First(&todo, userId)
	if err := res.Error; err != nil {
		return nil
	}

	todo.Title = payload.Title
	todo.Completed = payload.Completed
	repository.DB.Save(todo)

	return todo
}

func (repository TodoRepository) Delete(id uint) *entity.TodoEntity {
	var todo *entity.TodoEntity

	if err := repository.DB.First(&todo, id).Error; err != nil {
		return nil
	}
	repository.DB.Delete(todo)

	return todo
}
