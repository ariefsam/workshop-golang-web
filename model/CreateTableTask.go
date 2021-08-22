package model

import (
	"database/sql"
	"fmt"
)

func CreateTableTask() (err error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", MySQL.User, MySQL.Password, MySQL.Database))
	if err != nil {
		return
	}

	sqlStmt := `
	   CREATE TABLE if not exists task (
		   id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
		   name VARCHAR(255) DEFAULT "",
		   description TEXT,
		   status VARCHAR(255) DEFAULT "",
		   owner_id INTEGER DEFAULT 0
	   );
	`

	_, err = db.Exec(sqlStmt)
	db.Close()
	return
}
