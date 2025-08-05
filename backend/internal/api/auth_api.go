package api

import (
	"time"

	"github.com/SumirVats2003/formify/backend/internal/dbconnector"
	"github.com/SumirVats2003/formify/backend/internal/models"
	"github.com/SumirVats2003/formify/backend/utils"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(utils.GetEnv("JWT_SECRET", ""))

type UserApi struct {
	db            *mongo.Database
	userConnector dbconnector.UserDBConnector
}

func InitUserApi(db *mongo.Database) UserApi {
	u := dbconnector.InitUserConnector(db)
	userApi := UserApi{db: db, userConnector: u}
	return userApi
}

func (u UserApi) LoginApi(loginRequest models.LoginRequest) (string, error) {
	passwordHash, err := hashPassword(loginRequest.Password)
	if err != nil {
		return "", err
	}

	isAuthenticated, err := u.userConnector.LoginUser(loginRequest.Email, passwordHash)
	if err != nil {
		return "", err
	}
	if !isAuthenticated {
		return "", nil
	}

	userToken, err := createToken(loginRequest.Email)
	if err != nil {
		return "", err
	}

	return userToken, nil
}

func (u UserApi) Signup(signupRequest models.SignupRequest) (string, error) {
	hashedPassword, err := hashPassword(signupRequest.Password)
	if err != nil {
		return "", err
	}

	signupRequest.Password = hashedPassword
	u.userConnector.SignupUser(signupRequest)
	return "", nil
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
