package api

import (
	"time"

	"github.com/SumirVats2003/formify/backend/internal/database"
	"github.com/SumirVats2003/formify/backend/internal/models"
	"github.com/SumirVats2003/formify/backend/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(utils.GetEnv("JWT_SECRET", ""))

func LoginApi(loginRequest models.LoginRequest) (string, error) {
	passwordHash := hashPassword(loginRequest)

	isAuthenticated, err := database.LoginUser(loginRequest.Email, passwordHash)
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

func Signup(signupRequest models.SignupRequest) (string, error) {
	// return signup token
	return "", nil
}

func createToken(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"iss": "todo-app",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func hashPassword(loginRequest models.LoginRequest) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginRequest.Password), 14)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}
