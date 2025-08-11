package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/SumirVats2003/formify/backend/internal/api"
	"github.com/SumirVats2003/formify/backend/internal/models"
	"github.com/SumirVats2003/formify/backend/utils"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FormRouter struct {
	db      *mongo.Database
	formApi api.FormApi
}

func InitFormRoutes(db *mongo.Database, ctx context.Context) chi.Router {
	formApi := api.InitFormApi(db, ctx)
	f := FormRouter{db: db, formApi: formApi}
	r := chi.NewRouter()
	r.Post("/user/{userId}/create-form", f.CreateForm)
	r.Get("/{formId}", f.GetFormById)
	r.Delete("/{formId}", f.DeleteFormById)
	r.Get("/user/{userId}/all-form-summaries", f.GetAllUserFormSummaries)
	return r
}

func (f FormRouter) CreateForm(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	if userId == "" {
		http.Error(w, "Missing Parameters", http.StatusBadRequest)
	}

	var formRequest models.FormRequest
	formRequest, err := utils.ParseJSON(formRequest, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	formId, err := f.formApi.CreateForm(userId, formRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"formId": formId,
	})
}

func (f FormRouter) GetFormById(w http.ResponseWriter, r *http.Request) {
	formId := chi.URLParam(r, "formId")
	if formId == "" {
		http.Error(w, "Missing Parameters", http.StatusBadRequest)
	}

	form, err := f.formApi.GetFormById(formId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]models.FormResponse{
		"form": form,
	})
}

func (f FormRouter) DeleteFormById(w http.ResponseWriter, r *http.Request) {
	formId := chi.URLParam(r, "formId")
	if formId == "" {
		http.Error(w, "Missing Parameters", http.StatusBadRequest)
	}

	deleted, err := f.formApi.DeleteFormById(formId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{
		"deleted": deleted,
	})
}

func (f FormRouter) GetAllUserFormSummaries(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	if userId == "" {
		http.Error(w, "Missing Parameters", http.StatusBadRequest)
	}

	formSummaries, err := f.formApi.GetAllUserFormSummaries(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]models.FormSummary{
		"form_summaries": formSummaries,
	})
}
