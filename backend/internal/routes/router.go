package routes

import (
	"github.com/SumirVats2003/formify/backend/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	// heartbeat
	r.Get("/heartbeat", app.Heartbeat)

	// route groups
	authRouter := InitAuthRoutes(app.DB)

	r.Mount("/auth", authRouter)
	return r
}
