package api

import (
	"context"

	"github.com/SumirVats2003/formify/backend/internal/repository"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FormApi struct {
	db             *mongo.Database
	formRepository repository.FormRepository
}

func InitFormApi(db *mongo.Database, ctx context.Context) (FormApi, error) {
	f := repository.InitFormRepository(db, ctx)
	formApi := FormApi{db: db, formRepository: f}
	return formApi, nil
}
