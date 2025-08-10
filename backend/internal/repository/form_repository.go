package repository

import (
	"context"

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

	_, err := coll.InsertOne(f.ctx, bson.M{
		"_id":  id,
		"form": form,
	})

	if err != nil {
		return "", err
	}

	return id.Hex(), nil
}
