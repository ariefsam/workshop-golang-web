package main

import (
	"log"
	"todo/model"
)

func main() {
	InitMySQL()
	err := model.CreateTableUser()
	if err != nil {
		log.Println(err)
	}
	err = model.CreateTableTask()
	if err != nil {
		log.Println(err)
	}
	Server()
}

func InitMySQL() {
	model.MySQL.Database = "todo"
	model.MySQL.Password = "password"
	model.MySQL.User = "root"
	model.MySQL.Host = "127.0.0.1"
}
