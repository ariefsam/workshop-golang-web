package main

import "todo/model"

func main() {
	InitMySQL()
	model.CreateTableUser()
	Server()
}

func InitMySQL() {
	model.MySQL.Database = "todo"
	model.MySQL.Password = "password"
	model.MySQL.User = "root"
}
