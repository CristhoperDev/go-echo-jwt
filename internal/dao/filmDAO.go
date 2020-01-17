package dao

import (
	"fmt"
	conn "go-echo-jwt/internal/connection"
	"go-echo-jwt/internal/model"
)

func InsertFilm(title string, description string) error {
	sql := "INSERT INTO film(title, description) VALUES (?, ?)"
	rows, err := conn.DbConn.Query(sql, title, description)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UpdateFilm(title string, description string, filmId int) error {
	sql := "UPDATE film SET title = ?, description = ? WHERE idfilm = ?"
	rows, err := conn.DbConn.Query(sql, title, description, filmId)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteFilm(filmId int) error {
	sql := "DELETE FROM film WHERE idfilm = ?"
	rows, err := conn.DbConn.Query(sql, filmId)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetAllFilms() ([]model.Film, error) {
	var result []model.Film
	sql := "SELECT * FROM film"
	rows, err := conn.DbConn.Query(sql)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	for rows.Next() {
		var row model.Film
		err = rows.Scan(&row.IdFilm, &row.Title, &row.Description, &row.CreatedAt, &row.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return result, err
		}
		result = append(result, row)
	}

	return result, nil
}

func GetFilm(filmId int) (model.Film, error) {
	var result model.Film
	sql := "SELECT * FROM film WHERE idfilm = ?"
	row := conn.DbConn.QueryRow(sql, filmId)
	err := row.Scan(&result.IdFilm, &result.Title, &result.Description, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	return result, nil
}