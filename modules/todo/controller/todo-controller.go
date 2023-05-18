package controller

import (
	"akbariskndr/todo-service-gin/modules/todo/entity"
	"akbariskndr/todo-service-gin/modules/todo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	Service *service.TodoService
}

func (controller TodoController) FindAll(c *gin.Context) {
	todos := controller.Service.FindAll()

	c.IndentedJSON(http.StatusOK, gin.H{"data": todos})
}

func (controller TodoController) Create(c *gin.Context) {
	var payload entity.TodoEntity

	if err := c.BindJSON(&payload); err != nil {
		return
	}

	todo := controller.Service.Create(payload)

	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

func (controller TodoController) FindOne(c *gin.Context) {
	id := c.Param("id")

	todo := controller.Service.FindOne(id)

	if todo == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": todo})
}

func (controller TodoController) Update(c *gin.Context) {
	id := c.Param("id")

	var payload entity.TodoEntity

	if err := c.BindJSON(&payload); err != nil {
		return
	}

	todo := controller.Service.Update(id, payload)

	if todo == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": todo})
}

func (controller TodoController) Delete(c *gin.Context) {
	id := c.Param("id")

	todo := controller.Service.Delete(id)

	if todo == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": todo})
}
