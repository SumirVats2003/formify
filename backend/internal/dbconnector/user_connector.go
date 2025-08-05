package dbconnector

import (
	"github.com/SumirVats2003/formify/backend/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserDBConnector struct {
	db *mongo.Database
}

func InitUserConnector(db *mongo.Database) UserDBConnector {
	u := UserDBConnector{db: db}
	return u
}

func (u UserDBConnector) LoginUser(email, passwordHash string) (bool, error) {
	return false, nil
}

func (u UserDBConnector) SignupUser(signupRequest models.SignupRequest) (bool, error) {
	return false, nil
}
