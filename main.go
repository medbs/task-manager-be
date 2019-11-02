package main

import (
	"fmt"
	"log"
	"net/http"
	//"strconv"
	"context"
	"time"
	"github.com/gin-gonic/gin"
	"taskmanager/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	Id        int  `json:"Field Int"`
	Title     string `json:"Field Str"`
}

var todosCollection *mongo.Collection


func init() {
	 client := config.Connect()
   db := client.Database("blog")
	 todosCollection = db.Collection("posts")
	}

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", createTodo)
		v1.GET("/", fetchAllTodo)
	//	v1.GET("/:id", fetchSingleTodo)
	//	v1.PUT("/:id", updateTodo)
	//	v1.DELETE("/:id", deleteTodo)
	}

	router.Run()
}

func createTodo(cxt *gin.Context) {

	title := cxt.PostForm("Title")
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

	todo := Todo{1, title}
	// fmt.Println(todosCollection)
	 fmt.Println(todo)
	_ , err := todosCollection.InsertOne(ct, todo)
	if err != nil {
		log.Fatal(err)
	}

	cxt.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "todo item created successfully",
	})
}

func fetchAllTodo(cxt *gin.Context) {

	var todos []*Todo
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	 //err := todosCollection.Find(ct,nil).All(todos)
	 cur,err := todosCollection.Find(ct, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ct) {
    // create a value into which the single document can be decoded
    var todo Todo
    err := cur.Decode(&todo)
    if err != nil {
        log.Fatal(err)
    }

    todos = append(todos, &todo)
}

	if len(todos) <= 0 {
		cxt.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "no todo found",
		})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   todos,
	})
}
