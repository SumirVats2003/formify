package db

import (
	"log"

	"github.com/SumirVats2003/formify/backend/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDB() (*mongo.Client, error) {
	uri := utils.GetEnv("MONGO_URI", "")

	if uri == "" {
		log.Fatal("Set your 'MONGO_URI' environment variable.")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	client.Database("formify")
	log.Println("Connected to database...")

	return client, nil
}
