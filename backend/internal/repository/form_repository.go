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
	filter := bson.D{{"id", formId}}
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
