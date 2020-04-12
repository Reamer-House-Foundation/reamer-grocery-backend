package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getURL() string {
	password := os.Getenv("PASSWORD")
	query := "mongodb+srv://reamerdb:%s@reamergrocery-vr3dj.mongodb.net/test?retryWrites=true&w=majority"
	return fmt.Sprintf(query, password)
}

var query string = "{listing_url: \"https://www.airbnb.com/rooms/10006546\"}"

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI(getURL())

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("sample_airbnb").Collection("listingsAndReviews")

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

	fmt.Println(collection)
}
