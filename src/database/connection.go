package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func DbInstance() *mongo.Client {
	Uri := os.Getenv("MONGO_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(Uri))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb")
	Client = client
	return client

}

func GetClient() *mongo.Client {
	return Client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	if client == nil {
		DbInstance()
		return Client.Database("restaurant").Collection(collectionName)
	}
	return client.Database("restaurant").Collection(collectionName)
}
