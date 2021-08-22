package model_test

import (
	"log"
	"testing"
	"todo/model"
)

func TestListTask(t *testing.T) {
	InitMySQL()
	var tasks []model.Task
	var err error
	var userID int

	userID = 1

	tasks, err = model.ListTask(userID)
	if err != nil {
		t.Error(err.Error())
	}
	if len(tasks) == 0 {
		t.Error("Tasks kosong")
	}

	log.Println(tasks)

}
