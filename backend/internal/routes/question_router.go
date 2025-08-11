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

type QuestionRouter struct {
	db          *mongo.Database
	questionApi api.QuestionApi
}

func InitQuestionRoutes(db *mongo.Database, ctx context.Context) chi.Router {
	questionApi := api.InitQuestionApi(db, ctx)
	q := QuestionRouter{db: db, questionApi: questionApi}
	r := chi.NewRouter()
	r.Post("/form/{formId}/add", q.AddQuestion)
	r.Delete("/{questionId}/form/{formId}", q.DeleteQuestionById)
	return r
}

func (q QuestionRouter) AddQuestion(w http.ResponseWriter, r *http.Request) {
	formId := chi.URLParam(r, "formId")
	if formId == "" {
		http.Error(w, "Missing Parameters", http.StatusBadRequest)
	}

	var questionRequest models.QuestionRequest
	questionRequest, err := utils.ParseJSON(questionRequest, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	questionId, err := q.questionApi.AddQuestion(questionRequest, formId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"questionId": questionId,
	})
}

func (q QuestionRouter) DeleteQuestionById(w http.ResponseWriter, r *http.Request) {
	questionId := chi.URLParam(r, "questionId")
	formId := chi.URLParam(r, "formId")
	if formId == "" || questionId == "" {
		http.Error(w, "Missing Parameters", http.StatusBadRequest)
	}

	deleted, err := q.questionApi.DeleteQuestionById(questionId, formId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{
		"deleted": deleted,
	})
}
