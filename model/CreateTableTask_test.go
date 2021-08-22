package model_test

import (
	"testing"
	"todo/model"
)

func TestCreateTableTask(t *testing.T) {
	InitMySQL()
	err := model.CreateTableTask()
	if err != nil {
		t.Error(err.Error())
	}
}
