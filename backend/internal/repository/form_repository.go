package repository

import (
	"context"

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
