package model_test

import (
	"testing"
	"todo/model"
)

func TestEditTask(t *testing.T) {
	InitMySQL()
	taskID := 1
	ownerID := 1

	name := "Task diedt"
	description := "Description diedit"

	err := model.EditTask(taskID, ownerID, name, description)
	if err != nil {
		t.Error(err.Error())
	}

}
