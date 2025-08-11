package api

import (
	"context"

	"github.com/SumirVats2003/formify/backend/internal/models"
	"github.com/SumirVats2003/formify/backend/internal/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type QuestionApi struct {
	db                 *mongo.Database
	questionRepository repository.QuestionRepository
	formRepository     repository.FormRepository
}

func InitQuestionApi(db *mongo.Database, ctx context.Context) QuestionApi {
	q := repository.InitQuestionRepository(db, ctx)
	f := repository.InitFormRepository(db, ctx)
	questionApi := QuestionApi{db: db, questionRepository: q, formRepository: f}
	return questionApi
}

func (q QuestionApi) AddQuestion(question models.QuestionRequest, formId string) (string, error) {
	qId, err := q.questionRepository.CreateQuestion(question)
	if err != nil {
		return "", err
	}

	err = q.formRepository.AddQuestionToForm(formId, qId)
	if err != nil {
		return "", err
	}
	return qId, nil
}

func (q QuestionApi) DeleteQuestionById(questionId, formId string) (bool, error) {
	err := q.formRepository.RemoveQuestionFromForm(questionId, formId)
	if err != nil {
		return false, err
	}

	success, err := q.questionRepository.DeleteQuestionById(questionId)
	if err != nil {
		return false, err
	}
	return success, nil
}
