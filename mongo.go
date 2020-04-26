package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Grocery struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

func getGroceryByID(ID primitive.ObjectID, ctx context.Context,
	groceryCollection *mongo.Collection) (Grocery, error) {
	var result Grocery

	// Filter example
	err := groceryCollection.FindOne(ctx, bson.M{"_id": ID}).Decode(&result)
	if err != nil {
		fmt.Println("Failed getting grocery!")
		fmt.Println(err)
		return result, err
	}

	fmt.Println("Grocery: ", result)

	return result, nil
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("Not able to load enviornment file!")
	}
	return nil
}

func getURI() (string, error) {
	err := loadEnv()
	if err != nil {
		return "", err
	}

	return os.Getenv("MONGO_URI"), nil
}

/**
* As a user of this function you will have to disconnect from the database
 */
func connectToDB(uri string, dbname string) (*mongo.Database, context.Context) {
	if dbname == "" {
		log.Fatal("DBName not set in .env")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(dbname)

	return db, ctx
}

func getGroceryCollection(db *mongo.Database, name string) *mongo.Collection {
	return db.Collection(name)
}
