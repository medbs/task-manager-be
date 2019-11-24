package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"taskmanager/internal/config"
	"taskmanager/storage/model"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var tasksCollection *mongo.Collection

func init() {
	client := config.Connect()
	db := client.Database("tm")
	tasksCollection = db.Collection("tasks")
}

func CreateTask(cxt *gin.Context) {

	var task model.Task
	cxt.BindJSON(&task)
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newTask := model.Task{task.Title}
	fmt.Println(task)
	_, err := tasksCollection.InsertOne(ct, newTask)

	if err != nil {
		log.Fatal(err)
	}
	if task.Title == "" {
		cxt.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "task cannot have an empty title",
		})
	} else {
		cxt.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "task created successfully",
		})
	}
}

func UpdateTask(cxt *gin.Context) {

	var updatedTask model.Task
	title := cxt.Param("title")
	cxt.BindJSON(&updatedTask)

	updatedTitle := updatedTask.Title

	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := tasksCollection.UpdateOne(ct, bson.D{{"title", title}}, bson.M{"$set": bson.M{"title": updatedTitle}})

	if err != nil {
		fmt.Println("Error: " + err.Error())
		cxt.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "task not found",
		})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "task updated successfully!",
	})
}

func GetTask(cxt *gin.Context) {
	var task model.Task
	title := cxt.Param("title")
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := tasksCollection.FindOne(ct, bson.D{{"title", title}}).Decode(&task)

	if err != nil || task == (model.Task{}) {
		fmt.Println("Error: " + err.Error())
		cxt.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "task not found",
		})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   task,
	})
}

func GetTasks(cxt *gin.Context) {

	var tasks []*model.Task
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := tasksCollection.Find(ct, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ct) {
		// create a value into which the single document can be decoded
		var task model.Task
		err := cur.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, &task)
	}
	if len(tasks) <= 0 {
		cxt.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "no task found",
		})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   tasks,
	})
}

func DeleteTask(cxt *gin.Context) {

	title := cxt.Param("title")
	ct, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := tasksCollection.DeleteMany(ct, bson.D{{"title", title}})

	if err != nil {
		fmt.Println("Error: " + err.Error())
		cxt.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "task not found",
		})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "task with the name " + title + " is deleted successfully!",
	})
}
