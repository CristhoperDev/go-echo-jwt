package dao

import (
	"fmt"
	conn "go-echo-jwt/internal/connection"
	"go-echo-jwt/internal/model"
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

func UserLoginByUsernamePassword(username string, password string) ([]model.User, error) {
	var result []model.User
	sql := "SELECT * FROM user WHERE username = ? AND password = ?"
	rows, err := conn.DbConn.Query(sql, username, password)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	for rows.Next() {
		var row model.User
		err = rows.Scan(&row.IdUser, &row.Username, &row.Password, &row.CreatedAt, &row.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return result, err
		}
		result = append(result, row)
	}

	return result, nil
}

func getAllUsers() ([]model.User, error) {
	var result []model.User
	sql := "SELECT * FROM user"
	rows, err := conn.DbConn.Query(sql)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	for rows.Next() {
		var row model.User
		err = rows.Scan(&row.IdUser, &row.Username, &row.Password, &row.CreatedAt, &row.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return result, err
		}
		result = append(result, row)
	}

	return result, nil
}

