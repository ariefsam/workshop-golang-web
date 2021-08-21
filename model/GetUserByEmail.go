package model

import (
	"database/sql"
	"fmt"
)

func GetUserByEmail(email string) (user *User, err error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", MySQL.User, MySQL.Password, MySQL.Database))
	if err != nil {
		return
	}
	defer db.Close()

	stmt := `SELECT id,email,password FROM user WHERE email=?`
	rows, err := db.Query(stmt, email)
	defer rows.Close()

	for rows.Next() {
		var temp User
		err = rows.Scan(&temp.ID, &temp.Email, &temp.Password)
		user = &temp
	}
	return
}
