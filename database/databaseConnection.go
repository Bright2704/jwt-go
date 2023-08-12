package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	 "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func DBinstance() *mongo.Client{
	err := godotenv.Load(".env")
	if err !=nil{
		log.Fatal("Error loading .env file")
	}
	
	AuthSource := os.Getenv("AUTHSOURCE")
	//Mongodb := os.Getenv("MONGODB_URL")

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	clientOptions.Auth = &options.Credential{
		Username: os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		AuthSource: AuthSource,
		
		
	}
	

	// MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(os.Getenv("AUTHSOURCE")).Collection(collectionName)
		return collection
	}

