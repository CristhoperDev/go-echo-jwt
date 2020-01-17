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

func UpdateFilm(title string, description string, id_film int) error {
	sql := "UPDATE film SET title = ?, description = ? WHERE idfilm = ?"
	rows, err := conn.DbConn.Query(sql, title, description, id_film)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}