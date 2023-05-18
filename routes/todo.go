package routes

import (
	"akbariskndr/todo-service-gin/modules/todo/controller"

	"github.com/gin-gonic/gin"
)

func AddTodoRoutes(rg *gin.RouterGroup, controller *controller.TodoController) {
	router := rg.Group("/todos")

	router.GET("/", controller.FindAll)
	router.GET("/:id", controller.FindOne)
	router.POST("/", controller.Create)
	router.PATCH("/:id", controller.Update)
	router.DELETE("/:id", controller.Delete)
}
