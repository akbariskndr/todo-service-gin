package service

import (
	"akbariskndr/todo-service-gin/modules/auth/controller/dto"
	"akbariskndr/todo-service-gin/modules/auth/entity"
	"akbariskndr/todo-service-gin/modules/auth/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repository *repository.UserRepository
}

func (service AuthService) FindAll() []entity.UserEntity {
	return service.Repository.FindAll()
}

func (service AuthService) FindOne(id uint) *entity.UserEntity {
	return service.Repository.FindOne(id)
}

func (service AuthService) Register(payload *dto.RegisterDto) *entity.UserEntity {
	password := []byte(payload.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		return nil
	}

	payload.Password = string(hashedPassword)
	return service.Repository.Create(payload)
}

func (service AuthService) Login(payload *dto.LoginDto) *entity.UserEntity {
	user := service.Repository.FindByEmail(payload.Email)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil
	}

	return user
}

func (service AuthService) FindByEmail(email string) *entity.UserEntity {
	return service.Repository.FindByEmail(email)
}

func (service AuthService) Update(id string, payload entity.UserEntity) *entity.UserEntity {
	return service.Repository.Update(id, payload)
}

func (service AuthService) ChangePassword(id uint, payload entity.UserEntity) *entity.UserEntity {
	user := service.Repository.FindOne(id)

	if user == nil {
		return nil
	}

	password := []byte(payload.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	service.Repository.ChangePassword(id, string(hashedPassword))

	return user
}

func (service AuthService) Delete(id string) *entity.UserEntity {
	return service.Repository.Delete(id)
}
