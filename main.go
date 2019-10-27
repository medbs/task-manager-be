package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"context"
	"time"
	"github.com/gin-gonic/gin"
	"taskmanager/pkg/config"
//	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        int `json:"Field Int"`
	Title     string `json:"Field Str"`
	Completed bool `json:"Field Bool"`
}

var todosCollection *mongo.Collection


func init() {
	 client := config.Connect()
   db := client.Database("blog")
	 todosCollection := db.Collection("posts")
	 _ = todosCollection
	}

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", createTodo)
	//	v1.GET("/", fetchAllTodo)
	//	v1.GET("/:id", fetchSingleTodo)
	//	v1.PUT("/:id", updateTodo)
	//	v1.DELETE("/:id", deleteTodo)
	}

	router.Run()
}

func createTodo(cxt *gin.Context) {

	//title := cxt.PostForm("Title")

	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

	//var todo = Todo{1, title, false}
	todo := Todo{1, "kek", true}
	fmt.Println("" + todo.Title + " completed: " + strconv.FormatBool(todo.Completed))
	_ , err := todosCollection.InsertOne(ct,&todo)
	if err != nil {
		log.Fatal(err)
	}

	cxt.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "todo item created successfully",
	})
}

/*func fetchAllTodo(cxt *gin.Context) {
	var todos []Todo
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ , err := todosCollection.Find(ct,nil).All(&todos)
	if err != nil {
		log.Fatal(err)
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
} */

/*	func fetchSingleTodo(cxt *gin.Context) {
	todo := Todo{}
	id := bson.ObjectIdHex(cxt.Param("id"))
	_ ,	err := todosCollection.FindId(id).One(&todo)

	if err != nil || todo == (Todo{}) {
		fmt.Println("Error: " + err.Error())
		cxt.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "todo not found",
		})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   todo,
	})
} */

/*func updateTodo(cxt *gin.Context) {
	id := bson.ObjectIdHex(cxt.Param("id"))
	title := cxt.PostForm("title")
	completed, _ := strconv.ParseBool(cxt.PostForm("completed"))

	_,err := todosCollection.UpdateOne(id, bson.M{"title": title, "completed": completed})

	fmt.Printf("completed: %t\n\n", completed)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		cxt.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "todo not found",
		})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo updated successfully!",
	})
} */

/* func deleteTodo(cxt *gin.Context) {
	id := bson.ObjectIdHex(cxt.Param("id"))

	fmt.Printf("id: %v", id)

	_,err := todosCollection.RemoveId(id)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		cxt.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "todo not found",
		})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deleted successfully!",
	})

} */
