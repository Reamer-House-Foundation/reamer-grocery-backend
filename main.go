package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	uri, err := getURI()
	if err != nil {
		log.Fatal()
	}
	fmt.Println("URI: ", uri)

	// Getting the db name from the env file, is this correct?
	dbname := os.Getenv("DBNAME")
	db, ctx := connectToDB(uri, dbname)

	fmt.Println(db, ctx)

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
