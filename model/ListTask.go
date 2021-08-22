package model

import (
	"database/sql"
	"fmt"
)

func ListTask(userID int) (tasks []Task, err error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", MySQL.User, MySQL.Password, MySQL.Database))
	if err != nil {
		return
	}
	defer db.Close()

	stmt := `SELECT id, name, description, status, owner_id FROM task WHERE owner_id=?`
	rows, err := db.Query(stmt, userID)
	defer rows.Close()

	for rows.Next() {
		var temp Task
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Description, &temp.Status, &temp.OwnerID)
		tasks = append(tasks, temp)
	}
	return
}
