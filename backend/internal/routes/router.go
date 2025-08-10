package routes

import (
	"github.com/SumirVats2003/formify/backend/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.App) (*chi.Mux, error) {
	r := chi.NewRouter()
	db := app.DB.Database("formify")

	// heartbeat
	r.Get("/heartbeat", app.Heartbeat)

	// route groups
	authRouter, err := InitAuthRoutes(db, app.Ctx)
	if err != nil {
		return nil, err
	}
	formRouter, err := InitFormRoutes(db, app.Ctx)

	r.Mount("/auth", authRouter)
	r.Mount("/form", formRouter)

	return r, nil
}
