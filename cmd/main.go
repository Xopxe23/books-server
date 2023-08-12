package main

import (
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/xopxe23/books-server/internal/repository"
	"github.com/xopxe23/books-server/internal/service"
	"github.com/xopxe23/books-server/internal/transport/rest"
	"github.com/xopxe23/books-server/pkg/database"
)

// @title Books App API
// @version 1.0
// @description API Server for Books Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSL_MODE"),
	})
	if err != nil {
		logrus.Fatalf("error creating postgres connection: %s", err.Error())
	}
	repo := repository.NewRepostory(db)
	service := service.NewService(repo)
	handler := rest.NewHandler(service)

	srv := http.Server{
		Addr:    ":8000",
		Handler: handler.InitRoutes(),
	}
	if err = srv.ListenAndServe(); err != nil {
		logrus.Fatalf("error runing server: %s", err.Error())
	}
}
