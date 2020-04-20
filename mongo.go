package main

import (
	"errors"
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
	err := loadEnv()
	if err != nil {
		return "", err
	}

	return os.Getenv("MONGO_URI"), nil
}

func getDBName() (string, error) {
	dbname := os.Getenv("DBNAME")
	if dbname == "" {
		return "", errors.New("DBName not defined in env file!")
	}

	return dbname, nil
}
