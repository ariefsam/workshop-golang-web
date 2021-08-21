package model

import (
	"database/sql"
	"fmt"
)

func CreateUser(email, password string) (id int, err error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", MySQL.User, MySQL.Password, MySQL.Database))
	if err != nil {
		return
	}

	sqlStmt := `
	   INSERT INTO user(email, password) VALUES (?,?)
	`

	_, err = db.Exec(sqlStmt, email, password)
	db.Close()
	return
}
