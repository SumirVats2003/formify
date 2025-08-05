package dbconnector

import (
	"github.com/SumirVats2003/formify/backend/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthConnector struct {
	db *mongo.Database
}

func InitAuthConnector(db *mongo.Database) AuthConnector {
	a := AuthConnector{db: db}
	return a
}

func (a AuthConnector) LoginUser(email, passwordHash string) (bool, error) {
	return false, nil
}

func (a AuthConnector) SignupUser(signupRequest models.SignupRequest) (bool, error) {
	return false, nil
}
