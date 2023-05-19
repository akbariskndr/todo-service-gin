package service

import (
	"akbariskndr/todo-service-gin/lib/mailer"
	"akbariskndr/todo-service-gin/lib/view"
	"akbariskndr/todo-service-gin/modules/auth/controller/dto"
	"akbariskndr/todo-service-gin/modules/auth/entity"
	"akbariskndr/todo-service-gin/modules/auth/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repository *repository.UserRepository
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

	if user == nil {
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil
	}

	return user
}

func (service AuthService) ChangePassword(id uint, payload *dto.ChangePasswordDto) *entity.UserEntity {
	user := service.Repository.FindOne(id)

	if user == nil {
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil
	}

	password := []byte(payload.NewPassword)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	service.Repository.ChangePassword(id, string(hashedPassword))

	return user
}

func (service AuthService) ForgotPassword(email string) *entity.UserEntity {
	user := service.Repository.FindByEmail(email)

	if user == nil {
		return nil
	}

	to := []string{user.Email}
	subject := "Reset Password"

	view := view.Get("forgot-password.html")

	mailer.CreateBuilder().
		To(&to).
		Subject(&subject).
		Html(view.GetHtml()).
		Send()

	return user
}
