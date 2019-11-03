package rest

import (
	"fmt"
	"log"
	"net/http"
	"context"
	"time"
	"github.com/gin-gonic/gin"
	"taskmanager/storage/model"
	"taskmanager/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var todosCollection *mongo.Collection

func init(){
	client := config.Connect()
	db := client.Database("blog")
	todosCollection = db.Collection("posts")
}

func CreateTodo(cxt *gin.Context) {

	title := cxt.PostForm("title")
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

	todo := model.Todo{1, title}
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

func GetAllTodos(cxt *gin.Context) {

	var todos []*model.Todo
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur,err := todosCollection.Find(ct, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ct) {
    // create a value into which the single document can be decoded
    var todo model.Todo
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


func DeleteTodo(cxt *gin.Context) {

 title:= cxt.Param("title")
 ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
 defer cancel()
 _,err := todosCollection.DeleteMany(ct,bson.D{{"title",title}})

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
	 "message": "task with the name "+ title + " is deleted successfully!",
 })
}


func GetSingleTodo(cxt *gin.Context) {
var todo model.Todo
title:= cxt.Param("title")
ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
err := todosCollection.FindOne(ct,bson.D{{"title",title}}).Decode(&todo)

if err != nil || todo == (model.Todo{}) {
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
}


func UpdateTodo(cxt *gin.Context) {
	title:= cxt.Param("title")
	updatedTitle := cxt.PostForm("updatedTitle")
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_,err := todosCollection.UpdateOne(ct,bson.D{{"title",title}},bson.M{"$set": bson.M{"title": updatedTitle}})
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
}
