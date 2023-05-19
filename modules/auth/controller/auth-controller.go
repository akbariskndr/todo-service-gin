package controller

import (
	"akbariskndr/todo-service-gin/modules/auth/controller/dto"
	"akbariskndr/todo-service-gin/modules/auth/entity"
	"akbariskndr/todo-service-gin/modules/auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthController struct {
	AuthService *service.AuthService
	UserService *service.UserService
}

func (controller AuthController) Register(c *gin.Context) {
	var payload dto.RegisterDto

	if err := c.ShouldBindWith(&payload, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if user := controller.UserService.FindByEmail(payload.Email); user != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User already registered with that email"})
		return
	}

	user := controller.AuthService.Register(&payload)

	c.IndentedJSON(http.StatusOK, gin.H{"data": user})
}

func (controller AuthController) Me(c *gin.Context) {
	user, _ := c.Get("id")

	c.IndentedJSON(http.StatusOK, gin.H{"data": user})
}

func (controller AuthController) ChangePassword(c *gin.Context) {
	var payload dto.ChangePasswordDto

	if err := c.ShouldBindWith(&payload, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user, _ := c.Get("id")
	if res := controller.AuthService.ChangePassword(user.(*entity.UserEntity).ID, &payload); res == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": user})
}

func (controller AuthController) ForgotPasword(c *gin.Context) {
	var payload dto.ForgotPasswordDto

	if err := c.ShouldBindWith(&payload, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result := controller.AuthService.ForgotPassword(payload.Email)

	if result == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Check your email"})
}
