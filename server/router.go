package server

import (
	"taskmanager/rest"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/api/v1/todos")

	{
		v1.POST("/", rest.CreateTodo)
		v1.GET("/", rest.GetAllTodos)
		v1.GET("/:title", rest.GetSingleTodo)
		v1.PUT("/:title", rest.UpdateTodo)
		v1.DELETE("/:title", rest.DeleteTodo)
	}

	return router
}
