package models

type LoginRequest struct {
	Id       string `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
}

type User struct {
	UserId string        `json:"userId"`
	User   SignupRequest `json:"user"`
}
