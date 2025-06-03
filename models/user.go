package models

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}
type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
