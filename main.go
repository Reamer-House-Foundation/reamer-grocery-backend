package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	gql "github.com/reamer-house-foundation/reamer-grocery-graphql/pkg/graphql"
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

	if gql.ImportJSONDataFromFile("data.json", &data) {
		http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
			result := gql.ExecuteQuery(r.URL.Query().Get("query"), schema)
			json.NewEncoder(w).Encode(result)
		})

		fmt.Println("Now server is running on port 8080")
		fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={user(id:\"1\"){name}}'")
		http.ListenAndServe(":8080", nil)
	}
}

/******************************************************************************
*	THIS IS ALL GRAPHQL STUFF THAT NEEDS TO BE RE FACTORED AND ORGANIZED FOR
*	THE PKG ARRANGMENT
******************************************************************************/

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data map[string]user

//   Create User object type with fields "id" and "name" by using GraphQLObjectTypeConfig:
//       - Name: name of object type
//       - Fields: a map of fields by using GraphQLFields
//   Setup type of field use GraphQLFieldConfig
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

//   Create Query object type with fields "user" has type [userType] by using GraphQLObjectTypeConfig:
//       - Name: name of object type
//       - Fields: a map of fields by using GraphQLFields
//   Setup type of field use GraphQLFieldConfig to define:
//       - Type: type of field
//       - Args: arguments to query with current field
//       - Resolve: function to query data using params from [Args] and return value with current type
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return data[idQuery], nil
					}
					return nil, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)
