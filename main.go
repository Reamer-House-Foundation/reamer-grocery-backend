package main

import (
	"fmt"
	"log"
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
}
