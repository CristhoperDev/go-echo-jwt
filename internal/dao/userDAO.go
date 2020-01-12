package dao

import (
	"fmt"
	conn "go-echo-jwt/internal/connection"
)

func InsertUser(username string, password string) error {
	sql := "INSERT INTO user(username, password) VALUES (?, ?)"
	rows, err := conn.DbConn.Query(sql, username, password)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

