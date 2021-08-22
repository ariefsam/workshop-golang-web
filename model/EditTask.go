package model

import (
	"database/sql"
	"fmt"
)

func EditTask(taskID, ownerID int, name string, description string) (err error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", MySQL.User, MySQL.Password, MySQL.Database))
	if err != nil {
		return
	}

	sqlStmt := `
	   UPDATE task SET name=?, description=? WHERE id=? AND owner_id=?
	`

	_, err = db.Exec(sqlStmt, name, description, taskID, ownerID)
	if err != nil {
		return
	}
	db.Close()
	return
}
