package model

import (
	"database/sql"
	"fmt"
)

func ChangeTaskStatus(taskID int, ownerID int, status string) (err error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", MySQL.User, MySQL.Password, MySQL.Database))
	if err != nil {
		return
	}

	sqlStmt := `
	   UPDATE task SET status=? WHERE id=? AND owner_id=?
	`

	_, err = db.Exec(sqlStmt, status, taskID, ownerID)
	if err != nil {
		return
	}
	// temp, err := result.RowsAffected()
	// taskID = int(temp)
	db.Close()
	return
}
