package repo

import (
	"fmt"

	"github.com/reamer-house-foundation/reamer-grocery-graphql/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func (db *DB) GetGrocerys() ([]models.Grocery, error) {
	var results []models.Grocery

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

func (db *DB) GetGroceryByID(ID string) (models.Grocery, error) {
	var result models.Grocery
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

func (db *DB) GetGroceryByQuantity(quantity int) ([]models.Grocery, error) {
	var results []models.Grocery

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
