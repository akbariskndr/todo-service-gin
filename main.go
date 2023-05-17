package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt string `json:"created_at"`
}

type TodoList []*Todo

var todos = TodoList{
	{ID: "1", Title: "Learn Basic Golang", Completed: false, CreatedAt: "2023-05-18"},
	{ID: "2", Title: "Implement REST API", Completed: false, CreatedAt: "2023-05-18"},
	{ID: "3", Title: "Test the API", Completed: false, CreatedAt: "2023-05-18"},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func postTodos(c *gin.Context) {
	var newTodo *Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func findOneTodo(c *gin.Context) {
	id := c.Param("id")

	for _, val := range todos {
		if val.ID == id {
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var payload Todo

	for _, val := range todos {
		if val.ID == id {
			c.BindJSON(&payload)

			val.Title = payload.Title
			val.Completed = payload.Completed
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	var newList TodoList
	var deletedTodo *Todo = nil

	for i, val := range todos {
		if val.ID == id {
			deletedTodo = val
			newList = append(todos[:i], todos[i+1:]...)
		}
	}
	if deletedTodo != nil {
		todos = newList
		c.IndentedJSON(http.StatusOK, deletedTodo)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", postTodos)
	router.GET("/todos/:id", findOneTodo)
	router.PATCH("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)

	router.Run("localhost:8080")
}
