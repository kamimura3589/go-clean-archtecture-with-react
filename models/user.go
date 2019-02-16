package model

type User struct {
	ID       int    `json:"id"`
	email    string `json:"email"`
	password string `json:"password"`
}
