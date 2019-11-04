package config

import (
	"context"
	"fmt"
  "os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Connect() *mongo.Client {

  var clientOptions *options.ClientOptions

	// create a new context
	ctx := context.Background()

 	mongoURL := os.Getenv("MONGODB_HOST")

	// Set client options
	if mongoURL != "" {
		clientOptions = options.Client().ApplyURI("mongodb://"+mongoURL+":27017")
	} else {
    clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	}
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

/*func Disconnect() *mongo.Client{
  err = client.Disconnect(context.TODO())

  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Println("Connection to MongoDB closed.")
  }
} */
