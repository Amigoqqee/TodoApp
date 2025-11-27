package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func Connect() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI not set")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	log.Println("Connected to MongoDB Atlas")
	return client
}

func SetCollection(client *mongo.Client) {
	Collection = client.Database("golang_db").Collection("todos")
}