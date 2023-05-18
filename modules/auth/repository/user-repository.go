package repository

import (
	"akbariskndr/todo-service-gin/modules/auth/controller/dto"
	"akbariskndr/todo-service-gin/modules/auth/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repository UserRepository) FindAll() []entity.UserEntity {
	var users []entity.UserEntity
	repository.DB.Find(&users)

	return users
}

func (repository UserRepository) Create(payload *dto.RegisterDto) *entity.UserEntity {
	user := &entity.UserEntity{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	repository.DB.Create(user)

	return user
}

func (repository UserRepository) FindOne(id uint) *entity.UserEntity {
	var user *entity.UserEntity

	if err := repository.DB.First(&user, id).Error; err != nil {
		return nil
	}

	return user
}

func (repository UserRepository) FindByEmail(email string) *entity.UserEntity {
	var user *entity.UserEntity

	if err := repository.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil
	}

	return user
}

func (repository UserRepository) Update(id string, payload entity.UserEntity) *entity.UserEntity {
	var user *entity.UserEntity

	if err := repository.DB.First(&user, id).Error; err != nil {
		return nil
	}

	user.Name = payload.Name
	user.Email = payload.Email
	repository.DB.Save(user)

	return user
}

func (repository UserRepository) ChangePassword(id uint, password string) *entity.UserEntity {
	var user *entity.UserEntity

	if err := repository.DB.First(&user, id).Error; err != nil {
		return nil
	}

	user.Password = password
	repository.DB.Save(user)

	return user
}

func (repository UserRepository) Delete(id string) *entity.UserEntity {
	var user *entity.UserEntity

	if err := repository.DB.First(&user, id).Error; err != nil {
		return nil
	}
	repository.DB.Delete(user)

	return user
}
