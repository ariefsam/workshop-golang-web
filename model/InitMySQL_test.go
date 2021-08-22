package model_test

import "todo/model"

func InitMySQL() {
	model.MySQL.Database = "todo_test"
	model.MySQL.Password = "password"
	model.MySQL.User = "root"
}
