package utils

import "go.mongodb.org/mongo-driver/v2/bson"

func GenerateNewMongoId() bson.ObjectID {
	return bson.NewObjectID()
}
