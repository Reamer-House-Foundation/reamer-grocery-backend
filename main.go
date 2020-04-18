package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal()
	}

	uri, err := getURI()
	if err != nil {
		log.Fatal()
	}

	dbname, err := getDBName()
	if err != nil {
		log.Fatal()
	}

	fmt.Println("URI: ", uri, "DBName: ", dbname)
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
