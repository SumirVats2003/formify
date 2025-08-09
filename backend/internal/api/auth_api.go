package api

import (
	"context"
	"errors"
	"time"

	"github.com/SumirVats2003/formify/backend/internal/dbconnector"
	"github.com/SumirVats2003/formify/backend/internal/models"
	"github.com/SumirVats2003/formify/backend/utils"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(utils.GetEnv("JWT_SECRET", ""))

type AuthApi struct {
	db            *mongo.Database
	authConnector dbconnector.AuthConnector
}

func InitAuthApi(db *mongo.Database, ctx context.Context) AuthApi {
	u := dbconnector.InitAuthConnector(db, ctx)
	userApi := AuthApi{db: db, authConnector: u}
	return userApi
}

func (a AuthApi) Login(loginRequest models.LoginRequest) (string, error) {
	databaseDocument := a.authConnector.LoginUser(loginRequest.Email)

	var user models.SignupRequest
	err := databaseDocument.Decode(&user)
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return "", err
	}

	userToken, err := createToken(loginRequest.Email)
	if err != nil {
		return "", err
	}

	return userToken, nil
}

func (a AuthApi) Signup(signupRequest models.SignupRequest) (models.User, error) {
	hashedPassword, err := hashPassword(signupRequest.Password)
	if err != nil {
		return models.User{}, err
	}

	signupRequest.Password = hashedPassword
	signupRequest.CreatedAt = time.Now().UnixMilli()
	success, err := a.authConnector.SignupUser(signupRequest)
	return success, err
}

func createToken(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"iss": "formify",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
