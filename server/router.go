package server

import (
	"taskmanager/rest"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/api/v1/tm")

	{
		v1.POST("/tasks", rest.CreateTask)
		v1.GET("/tasks", rest.GetTasks)
		v1.GET("/tasks/:title", rest.GetTask)
		v1.PUT("/tasks/:title", rest.UpdateTask)
		v1.DELETE("/tasks/:title", rest.DeleteTask)
	}

	return router
}
