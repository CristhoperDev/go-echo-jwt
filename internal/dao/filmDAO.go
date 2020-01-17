package dao

import (
	"fmt"
	conn "go-echo-jwt/internal/connection"
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