package routes

import (
	todos "akbariskndr/todo-service-gin/services"

	"github.com/gin-gonic/gin"
)

func AddTodosRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/todos")

	router.GET("/", todos.FindAll)
	router.GET("/:id", todos.FindOne)
	router.POST("/", todos.Create)
	router.PATCH("/:id", todos.Update)
	router.DELETE("/:id", todos.Delete)
}
