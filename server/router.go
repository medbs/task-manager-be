package server

import (
	"taskmanager/rest"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/api/v1/tm")

	{
		v1.POST("/", rest.CreateTask)
		v1.GET("/", rest.GetTasks)
		v1.GET("/:title", rest.GetTask)
		v1.PUT("/:title", rest.UpdateTask)
		v1.DELETE("/:title", rest.DeleteTask)
	}

	return router
}
