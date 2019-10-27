package config

import (
  "log"
  "fmt"
  "context"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)




func Connect() *mongo.Client{

  // create a new context
  ctx := context.Background()

  // Set client options
  clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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
