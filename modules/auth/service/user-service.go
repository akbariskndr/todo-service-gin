package service

import (
	"akbariskndr/todo-service-gin/modules/auth/entity"
	"akbariskndr/todo-service-gin/modules/auth/repository"
)

type UserService struct {
	Repository *repository.UserRepository
}

func (service UserService) FindAll() []entity.UserEntity {
	return service.Repository.FindAll()
}

func (service UserService) FindOne(id uint) *entity.UserEntity {
	return service.Repository.FindOne(id)
}

func (service UserService) FindByEmail(email string) *entity.UserEntity {
	return service.Repository.FindByEmail(email)
}

func (service UserService) Update(id string, payload entity.UserEntity) *entity.UserEntity {
	return service.Repository.Update(id, payload)
}

func (service UserService) Delete(id string) *entity.UserEntity {
	return service.Repository.Delete(id)
}
