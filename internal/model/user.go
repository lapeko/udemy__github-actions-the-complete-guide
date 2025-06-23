package model

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
