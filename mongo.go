package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("Not able to load enviornment file!")
	}
	return nil
}

func getURI() (string, error) {
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	if username == "" {
		return "", errors.New("Username not defined in env file!")

	} else if password == "" {
		return "", errors.New("Password not defined in env file!")
	}

	uri := "mongodb+srv://%s:%s@reamergrocery-vr3dj.mongodb.net/test?retryWrites=true&w=majority"

	return fmt.Sprintf(uri, username, password), nil
}

func getDBName() (string, error) {
	dbname := os.Getenv("DBNAME")
	if dbname == "" {
		return "", errors.New("DBName not defined in env file!")
	}

	return dbname, nil

}
