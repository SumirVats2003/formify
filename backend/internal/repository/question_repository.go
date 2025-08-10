package repository

import (
	"context"

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

	_, err := coll.InsertOne(q.ctx, bson.M{
		"_id":        id,
		"id": id.Hex(),
		"question":   question,
	})

	if err != nil {
		return "", err
	}
	return id.Hex(), nil
}
