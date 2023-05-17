package todos

import (
	"akbariskndr/todo-service-gin/database"
	"akbariskndr/todo-service-gin/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	var todos []entity.Todo
	database.Connector.Find(&todos)

	c.IndentedJSON(http.StatusOK, gin.H{"data": todos})
}

func Create(c *gin.Context) {
	var todo entity.Todo

	if err := c.BindJSON(&todo); err != nil {
		return
	}

	database.Connector.Create(&todo)

	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

func FindOne(c *gin.Context) {
	id := c.Param("id")

	var todo entity.Todo

	if err := database.Connector.First(&todo, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": todo})
}

func Update(c *gin.Context) {
	id := c.Param("id")

	var todo entity.Todo
	var payload entity.Todo

	if err := c.BindJSON(&payload); err != nil {
		return
	}

	if err := database.Connector.First(&todo, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	todo.Title = payload.Title
	todo.Completed = payload.Completed
	database.Connector.Save(todo)

	c.IndentedJSON(http.StatusOK, gin.H{"data": todo})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	var todo entity.Todo

	if err := database.Connector.First(&todo, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	database.Connector.Delete(todo)

	c.IndentedJSON(http.StatusOK, gin.H{"data": todo})
}
