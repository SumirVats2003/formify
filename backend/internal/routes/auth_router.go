package routes

import (
	"net/http"

	"github.com/SumirVats2003/formify/backend/internal/api"
	"github.com/SumirVats2003/formify/backend/internal/models"
	"github.com/SumirVats2003/formify/backend/utils"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthRouter struct {
	db *mongo.Database
	a  api.AuthApi
}

func InitAuthRoutes(db *mongo.Database) chi.Router {
	authApi := api.InitAuthApi(db)
	a := AuthRouter{db: db, a: authApi}
	r := chi.NewRouter()
	r.Get("/login", a.Login)
	r.Post("/signup", a.Signup)
	return r
}

func (a AuthRouter) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest
	loginRequest, err := utils.ParseJSON(loginRequest, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// redirect to api and create a database layer
}

func (a AuthRouter) Signup(w http.ResponseWriter, r *http.Request) {
	var signupRequest models.SignupRequest
	signupRequest, err := utils.ParseJSON(signupRequest, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
