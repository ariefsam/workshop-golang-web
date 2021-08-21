package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var MySQL struct {
	User     string
	Password string
	Database string
}

func CreateTableUser() (err error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", MySQL.User, MySQL.Password, MySQL.Database))
	if err != nil {
		return
	}

	sqlStmt := `
	   CREATE TABLE if not exists user (
		   id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
		   email TEXT,
		   password TEXT
	   );
	`

	_, err = db.Exec(sqlStmt)
	db.Close()
	return
}
