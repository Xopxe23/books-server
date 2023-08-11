package main

import (
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/xopxe23/books-server/pkg/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	_, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("SSL_MODE"),
	})
	if err != nil {
		logrus.Fatalf("error creating postgres connection: %s", err.Error())
	}
}
