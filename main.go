package main

import (
	"github.com/gin-gonic/gin"
	"taskmanager/internal/config"
	"taskmanager/rest"
)

func init() {
	 client := config.Connect()
   db := client.Database("blog")
	 todosCollection = db.Collection("posts")
	}

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", rest.createTodo)
		v1.GET("/", rest.getAllTodos)
		v1.GET("/:title", rest.getSingleTodo)
		v1.PUT("/:title", rest.updateTodo)
		v1.DELETE("/:title", rest.deleteTodo)
	}

	router.Run()
}
