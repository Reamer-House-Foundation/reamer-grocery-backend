package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	repo "github.com/reamer-house-foundation/reamer-grocery-graphql/pkg/repository"
)

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("Not able to load enviornment file!")
	}
	return nil
}

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal()
	}

	db, err := repo.NewDB(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DBNAME"))
	if err != nil {
		log.Fatal()
	}

	/* This is just a test query */
	id := "5ea5e2365cfb870a298bb36e"
	grocery, err := db.GetGroceryByID(id)
	if err != nil {
		log.Fatal()
	}

	fmt.Println(grocery)

	grocerys, err := db.GetGrocerys()
	if err != nil {
		log.Fatal()
	}

	fmt.Println(grocerys)

	if importJSONDataFromFile("data.json", &data) {
		http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
			result := executeQuery(r.URL.Query().Get("query"), schema)
			json.NewEncoder(w).Encode(result)
		})

		fmt.Println("Now server is running on port 8080")
		fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={user(id:\"1\"){name}}'")
		http.ListenAndServe(":8080", nil)
	}
}
