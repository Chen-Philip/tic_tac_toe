package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func DBinstance() *mongo.Client {
	// Loads environment variables from the .env file and returns an error if one occurs
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// os.Getenv("MONGODB_URL") reads the environmental variable "MONGODB_URL" and returns it as a string
	MongoDb := os.Getenv("MONGODB_URL")

	// Set a timeout context for the connection (context is a messenger, so in this case
	// it lets the API know to cancel after 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // cleans up resoucrese immediatly after function finished, even if it doesn't timeout

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("cluster0").Collection(collectionName)
}
