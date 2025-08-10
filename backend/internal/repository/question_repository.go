package repository

import (
	"context"
	"errors"

	"github.com/SumirVats2003/formify/backend/internal/models"
	"github.com/SumirVats2003/formify/backend/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type QuestionRepository struct {
	db             *mongo.Database
	ctx            context.Context
	collectionName string
}

func InitQuestionRepository(db *mongo.Database, ctx context.Context) QuestionRepository {
	q := QuestionRepository{
		db:             db,
		ctx:            ctx,
		collectionName: "questions",
	}
	return q
}

func (q QuestionRepository) CreateQuestion(question models.QuestionRequest) (string, error) {
	id := utils.GenerateNewMongoId()
	coll := q.db.Collection(q.collectionName)

	questionModel := models.Question{
		Id:         id.Hex(),
		Title:      question.Title,
		AnswerType: question.AnswerType,
		Required:   question.Required,
		Options:    question.Options,
	}

	_, err := coll.InsertOne(q.ctx, questionModel)

	if err != nil {
		return "", err
	}
	return id.Hex(), nil
}

func (q QuestionRepository) GetQuestionById(questionId string) (models.Question, error) {
	filter := bson.D{{"id", questionId}}
	document := q.db.Collection(q.collectionName).FindOne(q.ctx, filter)

	if document == nil {
		return models.Question{}, errors.New("Question Not Found")
	}

	var question models.Question
	err := document.Decode(&question)

	if err != nil {
		return models.Question{}, err
	}
	return question, nil
}
