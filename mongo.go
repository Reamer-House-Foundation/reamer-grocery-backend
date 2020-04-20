package main

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(dbname)

	return db, ctx
}
