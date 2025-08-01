package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SumirVats2003/formify/backend/db"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	Logger *log.Logger
	DB     *mongo.Database
}

func InitApp() (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	dbclient, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := dbclient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	app := &App{
		Logger: logger,
		DB:     dbclient.Database("formify"),
	}
	return app, nil
}

func (a *App) Heartbeat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}
