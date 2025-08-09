package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SumirVats2003/formify/backend/internal/app"
	"github.com/SumirVats2003/formify/backend/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using empty variables")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := ":8080"
	app, err := app.InitApp(ctx)
	if err != nil {
		panic(err)
	}

	r, err := routes.SetupRoutes(app)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("%s", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("server running on port :8080\n")

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}

	defer app.DB.Disconnect(ctx)
}
