package model

import (
	"database/sql"
	"fmt"
)

func CreateTask(task Task) (taskID int, err error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", MySQL.User, MySQL.Password, MySQL.Database))
	if err != nil {
		return
	}

	sqlStmt := `
	   INSERT INTO task(name, description, status, owner_id) VALUES (?,?,?,?)
	`

	result, err := db.Exec(sqlStmt, task.Name, task.Description, task.Status, task.OwnerID)
	if err != nil {
		return
	}
	temp, err := result.LastInsertId()
	taskID = int(temp)
	db.Close()
	return
}
