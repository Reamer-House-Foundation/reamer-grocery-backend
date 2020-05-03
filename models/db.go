package models

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GroceryDB interface {
	getGroceryByID(string) (Grocery, error)
	getGrocerys() ([]Grocery, error)
	getGrocerysByQuantity(int) ([]Grocery, error)
	addGrocery([]Grocery) error
}

type DB struct {
	db  *mongo.Database
	ctx context.Context
}

/**
* Connect to the given mongo database and return a DB struct to interact with
* that db
 */
func NewDB(uri string, dbname string) (*DB, error) {
	if dbname == "" {
		return nil, errors.New("DB Name was not given to NewDB function!")
	}
	if uri == "" {
		return nil, errors.New("URI was not given to NewDB function!")
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

	return &DB{
		db:  client.Database(dbname),
		ctx: ctx,
	}, nil
}
