package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func DBSet() *mongo.Client {
	client, err := mongo.NewClient(Options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.fatal(err)
	}

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = Client.ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongodb:( ")
		return nil
	}
	fmt.Println("Successfully connected to mongodb ")
	return Client

}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return collection

}
func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return productCollection

}
