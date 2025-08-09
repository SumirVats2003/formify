package dbconnector

import (
	"context"
	"fmt"

	"github.com/SumirVats2003/formify/backend/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthConnector struct {
	db             *mongo.Database
	collectionName string
	ctx            context.Context
}

func InitAuthConnector(db *mongo.Database, ctx context.Context) AuthConnector {
	a := AuthConnector{db: db, ctx: ctx}
	a.collectionName = "users"
	return a
}

func (a AuthConnector) LoginUser(email string) *mongo.SingleResult {
	filter := bson.D{{"email", email}}
	return a.db.Collection(a.collectionName).FindOne(a.ctx, filter)
}

func (a AuthConnector) SignupUser(signupRequest models.SignupRequest) (bool, error) {
	coll := a.db.Collection(a.collectionName)

	_, err := coll.InsertOne(a.ctx, signupRequest)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
