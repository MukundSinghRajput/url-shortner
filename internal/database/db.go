package db

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func init() {
	opts := options.Client().ApplyURI(os.Getenv("MONGO_DB"))

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal("failed to connect to the database")
	}

	fmt.Println("Connected to Mongo DB")
	Database = client.Database("url-shotner")
}
