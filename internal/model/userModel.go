package model

import "time"

type UserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	IdUser 		int 		`json:"id_user"`
	Username 	string 		`json:"username"`
	Password 	string 		`json:"password"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}