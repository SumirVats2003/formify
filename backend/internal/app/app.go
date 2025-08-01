package app

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	Logger *log.Logger
	DB     *mongo.Database
}

func InitApp() (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &App{
		Logger: logger,
		DB:     nil,
	}
	return app, nil
}
