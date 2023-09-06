package models

import (
	"database/sql"
	"fmt"
	"new-project2/db"
	"new-project2/helpers"
)

type User struct {
	Id	int `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateConn()

	sqlStatement := "SELECT * FROM users WHERE username = $1"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("query error")
		return false, err
	}

	match, err := helpers.ComparePassword(password, pwd)
	if! match {
		fmt.Println("Hash and password doesn't match")
		return false, err
	}

	return true, nil

}