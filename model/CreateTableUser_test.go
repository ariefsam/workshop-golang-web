package model_test

import (
	"testing"
	"todo/model"
)

func TestCreateTableUser(t *testing.T) {
	InitMySQL()
	model.CreateTableUser()
}

func InitMySQL() {
	model.MySQL.Database = "todo"
	model.MySQL.Password = "password"
	model.MySQL.User = "root"
}
