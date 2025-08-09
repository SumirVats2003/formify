package repository

import (
	"context"
	"errors"

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
	return findUser(a.db, a.ctx, email)
}

func (a AuthConnector) SignupUser(signupRequest models.SignupRequest) (models.User, error) {
	id := bson.NewObjectID()
	coll := a.db.Collection(a.collectionName)

	var existingUser struct {
		Email string `bson:"email"`
	}
	err := coll.FindOne(a.ctx, bson.M{"user.email": signupRequest.Email}).Decode(&existingUser)
	if err == nil {
		return models.User{}, errors.New("user already exists")
	} else if err != mongo.ErrNoDocuments {
		return models.User{}, err
	}

	_, err = coll.InsertOne(a.ctx, bson.M{"_id": id, "userId": id.Hex(), "user": signupRequest})
	if err != nil {
		return models.User{}, err
	}

	userDoc := models.User{
		UserId: id.Hex(),
		User:   signupRequest,
	}
	return userDoc, nil
}

func findUser(db *mongo.Database, ctx context.Context, email string) *mongo.SingleResult {
	filter := bson.D{{"user.email", email}}
	return db.Collection("users").FindOne(ctx, filter)
}
