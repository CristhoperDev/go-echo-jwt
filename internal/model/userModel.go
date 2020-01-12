package model

type UserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	IdUser 		int 		`json:"id_user"`
	Username 	string 		`json:"username"`
	Password 	string 		`json:"password"`
	CreatedAt 	string 		`json:"created_at"`
	UpdatedAt 	string 		`json:"updated_at"`
}