package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        int    `json:"id"`
	Item      string `json:"item"`
	Completed string `json:"completed"`
}

var todos []todo = []todo{
	{ID: 1, Item: "Buy milk", Completed: "false"},
	{ID: 2, Item: "Buy eggs", Completed: "false"},
	{ID: 3, Item: "Buy bread", Completed: "false"},
}

func getTodos(context *gin.Context) {
	// Convert the todos variable to JSON
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodos(context *gin.Context) {
	var newTodo todo
	// Take whatever JSON  inside request body and bind it to newTodo
	// If there is an error, return the error
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	// Add the new todo to the slice.
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	var id string = context.Param("id")
	var strid int64
	strid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return
	}
	// Convert the id from string to int
	// If there is an error, return the error
	if err := context.BindJSON(&id); err != nil {
		return
	}
	// Loop through the todos slice and find the todo with the id
	for _, todo := range todos {

		if todo.ID == int(strid) {
			context.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodos)
	router.GET("/todos/:id", getTodo)
	router.Run("localhost:9090")
}
