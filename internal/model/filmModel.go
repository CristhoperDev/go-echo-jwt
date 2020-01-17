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
