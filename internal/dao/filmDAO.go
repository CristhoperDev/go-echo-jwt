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