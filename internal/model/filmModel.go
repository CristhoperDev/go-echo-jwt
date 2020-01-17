package model

type FilmModelPost struct {
	Title 			string 		`json:"title"`
	Description 	string 		`json:"description"`
}

type FilmModelPut struct {
	IdFilm			int 		`json:"id_film"`
	Title 			string 		`json:"title"`
	Description 	string 		`json:"description"`
}

type Film struct {
	IdFilm			int 		`json:"id_film"`
	Title 			string 		`json:"title"`
	Description 	string 		`json:"description"`
	CreatedAt 		string 		`json:"created_at"`
	UpdatedAt 		string 		`json:"updated_at"`
}