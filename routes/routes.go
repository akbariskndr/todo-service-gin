package routes

import (
	"akbariskndr/todo-service-gin/middleware"
	auth_module "akbariskndr/todo-service-gin/modules/auth"
	todo_module "akbariskndr/todo-service-gin/modules/todo"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	var router = gin.Default()

	v1 := router.Group("v1")
	AddPingRoutes(v1)

	todo := todo_module.GetInstance()
	auth := auth_module.GetInstance()

	authMiddleware := middleware.CreateAuthMiddleware()

	AddTodoRoutes(v1, todo.Controller, authMiddleware)
	AddAuthRoutes(v1, auth.Controller, authMiddleware)

	return router
}
