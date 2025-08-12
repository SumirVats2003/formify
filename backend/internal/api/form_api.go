package api

import (
	"context"

	"github.com/SumirVats2003/formify/backend/internal/models"
	"github.com/SumirVats2003/formify/backend/internal/repository"
	"github.com/SumirVats2003/formify/backend/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FormApi struct {
	db                 *mongo.Database
	formRepository     repository.FormRepository
	questionRepository repository.QuestionRepository
}

func InitFormApi(db *mongo.Database, ctx context.Context) FormApi {
	f := repository.InitFormRepository(db, ctx)
	q := repository.InitQuestionRepository(db, ctx)
	formApi := FormApi{db: db, formRepository: f, questionRepository: q}
	return formApi
}

func (f FormApi) CreateForm(userId string, formRequest models.FormRequest) (string, error) {
	questionIds := make([]string, 0)
	for _, question := range formRequest.Questions {
		qId, err := f.questionRepository.CreateQuestion(question)
		if err != nil {
			return "", err
		}
		questionIds = append(questionIds, qId)
	}

	form := models.Form{
		Id:                    "",
		Title:                 formRequest.Title,
		CreatorId:             formRequest.CreatorId,
		QuestionIds:           questionIds,
		AttachedSheet:         "",
		CreationTimestamp:     utils.GetCurrentTimestamp(),
		ModificationTimestamp: utils.GetCurrentTimestamp(),
		ValidityTimestamp:     formRequest.ValidityTimestamp,
	}

	formId, err := f.formRepository.CreateForm(form)
	if err != nil {
		return "", err
	}
	return formId, nil
}

func (f FormApi) GetFormById(formId string) (models.FormResponse, error) {
	form, err := f.formRepository.GetFormById(formId)
	if err != nil {
		return models.FormResponse{}, err
	}

	formQuestions := make([]models.Question, 0)
	for _, questionId := range form.QuestionIds {
		question, err := f.questionRepository.GetQuestionById(questionId)
		if err != nil {
			return models.FormResponse{}, err
		}
		formQuestions = append(formQuestions, question)
	}

	formResponse := models.FormResponse{
		Id:                    form.Id,
		Title:                 form.Title,
		CreatorId:             form.CreatorId,
		Questions:             formQuestions,
		AttachedSheet:         form.AttachedSheet,
		CreationTimestamp:     form.CreationTimestamp,
		ModificationTimestamp: form.ModificationTimestamp,
		ValidityTimestamp:     form.ValidityTimestamp,
	}
	return formResponse, nil
}

func (f FormApi) DeleteFormById(formId string) (bool, error) {
	questionIds, err := f.formRepository.GetFormQuestionIds(formId)
	if err != nil {
		return false, err
	}

	for _, questionId := range questionIds {
		_, err := f.questionRepository.DeleteQuestionById(questionId)
		if err != nil {
			return false, err
		}
	}

	deletionResult, err := f.formRepository.DeleteFormById(formId)
	if err != nil {
		return false, err
	}
	return deletionResult, nil
}

func (f FormApi) GetAllUserFormSummaries(userId string) ([]models.FormSummary, error) {
	formSummaries, err := f.formRepository.GetAllUserFormSummaries(userId)
	if err != nil {
		return nil, err
	}
	return formSummaries, nil
}
