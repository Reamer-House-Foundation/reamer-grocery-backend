package models

import (
	"errors"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

/* Repeated code I know.  I don't want to make a whole env.go module just for
* this little thing.  What should we do here?
 */
func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("Not able to load enviornment file!")
	}
	return nil
}

/* I Implemented this test like this because I'm not sure how to mock out this
* all or if it is even possible in our current arrangment.  This is better than
* nothing though.
 */
func test_GetGroceryByID(t *testing.T) {
	err := loadEnv()
	if err != nil {
		t.Error("Unable to Load Env File")
	}

	db, err := NewDB(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DBNAME"))
	if err != nil {
		t.Error("Unable to create DB")
	}

	// At the time of writing this test, this grocery ID exists in the DB
	_, err = db.GetGroceryByID("5ea5e2365cfb870a298bb36e")
	if err != nil {
		t.Error("Unable to obtain valid grocery")
	}

	// At the time of writing this test, this grocery ID does not exist in the DB
	_, err = db.GetGroceryByID("5ea5e2365cfb870a298bb36f")
	if err != nil {
		t.Error("Unable to obtain invalid grocery")
	}
}
