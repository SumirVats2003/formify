package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SumirVats2003/formify/backend/internal/app"
	"github.com/SumirVats2003/formify/backend/internal/routes"
)

func main() {
	port := ":8080"
	app, err := app.InitApp()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(app)
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
		app.Logger.Fatal()
	}
}
