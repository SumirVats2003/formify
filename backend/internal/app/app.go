package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SumirVats2003/formify/backend/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	Logger *log.Logger
	DB     *mongo.Client
	Ctx    context.Context
}

func InitApp(ctx context.Context) (*App, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	dbclient, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}

	app := &App{
		Logger: logger,
		DB:     dbclient,
		Ctx:    ctx,
	}
	return app, nil
}

func (a *App) Heartbeat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}
