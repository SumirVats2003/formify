package routes

import (
	"context"
	"net/http"

	"github.com/SumirVats2003/formify/backend/internal/api"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FormRouter struct {
	db      *mongo.Database
	formApi api.FormApi
}

func InitFormRoutes(db *mongo.Database, ctx context.Context) (chi.Router, error) {
	formApi, err := api.InitFormApi(db, ctx)

	if err != nil {
		return nil, err
	}

	f := FormRouter{db: db, formApi: formApi}
	r := chi.NewRouter()
	r.Get("/:formId", f.GetFormById)
	r.Get("/user/:userId/all-forms", f.GetAllUserForms)
	r.Post("/create-form", f.CreateForm)
	r.Delete("/:formId", f.DeleteFormById)
	return r, nil
}

func (f FormRouter) CreateForm(w http.ResponseWriter, r *http.Request)      {}
func (f FormRouter) GetFormById(w http.ResponseWriter, r *http.Request)     {}
func (f FormRouter) DeleteFormById(w http.ResponseWriter, r *http.Request)  {}
func (f FormRouter) GetAllUserForms(w http.ResponseWriter, r *http.Request) {}
