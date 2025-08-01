package routes

import (
	"fmt"
	"net/http"

	"github.com/SumirVats2003/formify/backend/internal/app"
	"github.com/SumirVats2003/formify/backend/internal/models"
	"github.com/SumirVats2003/formify/backend/utils"
	"github.com/go-chi/chi/v5"
)

type Authenticator struct {
	app *app.App
}

func InitAuthRoutes(app *app.App) chi.Router {
	a := Authenticator{app: app}
	r := chi.NewRouter()
	r.Get("/login", a.Login)
	r.Post("/signup", a.Signup)
	return r
}

func (a Authenticator) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest
	loginRequest, err := utils.ParseJSON(loginRequest, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.app.Logger.Println(loginRequest)
	// redirect to api and create a database layer
}

func (a Authenticator) Signup(w http.ResponseWriter, r *http.Request) {
	var signupRequest models.SignupRequest
	loginRequest, err := utils.ParseJSON(signupRequest, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.app.Logger.Println(loginRequest)
}
