package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// This is the structure of our data
type Grocery struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Quantity int32              `bson:"quantity"`
}

func (db *DB) GetGrocerys() ([]Grocery, error) {
	var results []Grocery

	groceryCursor, err := db.db.Collection("dev1.0").Find(db.ctx, bson.M{})
	if err != nil {
		fmt.Println("Failed getting groceries!")
		fmt.Println(err)
		return results, err
	}

	err = groceryCursor.All(db.ctx, &results)
	if err != nil {
		fmt.Println("Failed getting cursor!")
		fmt.Println(err)
		return results, err
	}

	return results, nil
}

func (db *DB) GetGroceryByID(ID string) (Grocery, error) {
	var result Grocery
	id, _ := primitive.ObjectIDFromHex(ID)

	/* Collection name is hardcoded here.. need to discuss with team
	*  We will probably have multiple collections.. how do we handle that? */
	err := db.db.Collection("dev1.0").FindOne(db.ctx, bson.M{"_id": id}).Decode(&result)

	if err != nil {
		fmt.Println("Failed getting grocery by Id!")
		fmt.Println(err)
		return result, err
	}

	return result, nil
}

func (db *DB) GetGroceriesByQuantity(quantity int) ([]Grocery, error) {
	var results []Grocery

	/* Collection name is hardcoded here.. need to discuss with team
	*  We will probably have multiple collections.. how do we handle that? */
	groceryCursor, err := db.db.Collection("dev1.0").Find(db.ctx, bson.M{"quantity": quantity})
	if err != nil {
		fmt.Println("Failed getting grocery by quantity!")
		fmt.Println(err)
		return results, err
	}

	/* Just load all of the entries into the results slice */
	err = groceryCursor.All(db.ctx, &results)
	if err != nil {
		fmt.Println("Failed getting cursor!")
		fmt.Println(err)
		return results, err
	}

	return results, nil
}
