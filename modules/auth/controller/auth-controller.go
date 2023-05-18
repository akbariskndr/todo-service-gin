package controller

import (
	"akbariskndr/todo-service-gin/modules/auth/controller/dto"
	"akbariskndr/todo-service-gin/modules/auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthController struct {
	Service *service.AuthService
}

func (controller AuthController) Register(c *gin.Context) {
	var payload dto.RegisterDto

	if err := c.ShouldBindWith(&payload, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if user := controller.Service.FindByEmail(payload.Email); user != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User already registered with that email"})
		return
	}

	user := controller.Service.Register(&payload)

	c.IndentedJSON(http.StatusOK, gin.H{"data": user})
}

func (controller AuthController) Me(c *gin.Context) {
	user, _ := c.Get("id")

	c.IndentedJSON(http.StatusOK, gin.H{"data": user})
}

// TODO: Implement change password
func (controller AuthController) ChangePassword(c *gin.Context) {
	users := controller.Service.FindAll()

	c.IndentedJSON(http.StatusOK, gin.H{"data": users})
}
