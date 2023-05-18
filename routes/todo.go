package routes

import (
	"akbariskndr/todo-service-gin/modules/todo/controller"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func AddTodoRoutes(rg *gin.RouterGroup, controller *controller.TodoController, authMiddleware *jwt.GinJWTMiddleware) {
	router := rg.Group("/todos")

	router.GET("/", authMiddleware.MiddlewareFunc(), controller.FindAll)
	router.GET("/:id", authMiddleware.MiddlewareFunc(), controller.FindOne)
	router.POST("/", authMiddleware.MiddlewareFunc(), controller.Create)
	router.PATCH("/:id", authMiddleware.MiddlewareFunc(), controller.Update)
	router.DELETE("/:id", authMiddleware.MiddlewareFunc(), controller.Delete)
}
