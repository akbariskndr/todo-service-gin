package controller

import (
	"akbariskndr/todo-service-gin/modules/auth/entity"
	"akbariskndr/todo-service-gin/modules/todo/controller/dto"
	"akbariskndr/todo-service-gin/modules/todo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type TodoController struct {
	Service *service.TodoService
}

func (controller TodoController) FindAll(c *gin.Context) {
	var payload dto.FindAllTodoDto
	user, _ := c.Get("id")

	if err := c.ShouldBindQuery(&payload); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	todos := controller.Service.FindAll(user.(*entity.UserEntity).ID, &payload)

	c.IndentedJSON(http.StatusOK, gin.H{"data": todos})
}

func (controller TodoController) Create(c *gin.Context) {
	var payload dto.CreateTodoDto
	user, _ := c.Get("id")

	if err := c.ShouldBindWith(&payload, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	todo := controller.Service.Create(user.(*entity.UserEntity).ID, &payload)

	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

func (controller TodoController) FindOne(c *gin.Context) {
	todoId, _ := strconv.Atoi(c.Param("id"))
	user, _ := c.Get("id")

	todo := controller.Service.FindOne(user.(*entity.UserEntity).ID, uint(todoId))

	if todo == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": todo})
}

func (controller TodoController) Update(c *gin.Context) {
	todoId, _ := strconv.Atoi(c.Param("id"))

	var payload dto.UpdateTodoDto
	user, _ := c.Get("id")

	if err := c.ShouldBindWith(&payload, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	todo := controller.Service.Update(uint(todoId), &payload, user.(*entity.UserEntity).ID)

	if todo == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": todo})
}

func (controller TodoController) Delete(c *gin.Context) {
	todoId, _ := strconv.Atoi(c.Param("id"))
	user, _ := c.Get("id")

	if todo := controller.Service.FindOne(user.(*entity.UserEntity).ID, uint(todoId)); todo == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	todo := controller.Service.Delete(uint(todoId))

	c.IndentedJSON(http.StatusOK, gin.H{"data": todo})
}
