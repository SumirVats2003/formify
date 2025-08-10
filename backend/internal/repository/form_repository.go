package repository

import (
	"context"
	"errors"

	"github.com/SumirVats2003/formify/backend/internal/models"
	"github.com/SumirVats2003/formify/backend/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FormRepository struct {
	db             *mongo.Database
	collectionName string
	ctx            context.Context
}

func InitFormRepository(db *mongo.Database, ctx context.Context) FormRepository {
	f := FormRepository{db: db, ctx: ctx}
	f.collectionName = "forms"
	return f
}

func (f FormRepository) CreateForm(form models.Form) (string, error) {
	id := utils.GenerateNewMongoId()
	coll := f.db.Collection(f.collectionName)
	form.Id = id.Hex()

	_, err := coll.InsertOne(f.ctx, form)
	if err != nil {
		return "", err
	}

	return id.Hex(), nil
}

func (f FormRepository) GetFormById(formId string) (models.Form, error) {
	filter := bson.D{{Key: "id", Value: formId}}
	document := f.db.Collection(f.collectionName).FindOne(f.ctx, filter)
	if document == nil {
		return models.Form{}, errors.New("Form Not Found")
	}

	var form models.Form
	err := document.Decode(&form)
	if err != nil {
		return models.Form{}, err
	}

	return form, nil
}

func (f FormRepository) GetFormQuestionIds(formId string) ([]string, error) {
	form, err := f.GetFormById(formId)
	questionIds := make([]string, 0)
	if err != nil {
		return questionIds, err
	}

	questionIds = form.QuestionIds
	return questionIds, err
}

func (f FormRepository) DeleteFormById(formId string) (bool, error) {
	filter := bson.D{{Key: "id", Value: formId}}
	_, err := f.db.Collection(f.collectionName).DeleteOne(f.ctx, filter)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (f FormRepository) GetAllUserFormSummaries(userId string) ([]models.FormSummary, error) {
	filter := bson.D{{Key: "creatorid", Value: userId}}
	cursor, err := f.db.Collection(f.collectionName).Find(f.ctx, filter)
	if err != nil {
		return nil, errors.New("Form Not Found")
	}

	var formSummaries []models.FormSummary
	err = cursor.All(f.ctx, &formSummaries)

	if err != nil {
		return nil, err
	}

	return formSummaries, nil
}
