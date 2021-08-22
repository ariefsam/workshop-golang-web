package main

import (
	"log"
	"os"
	"todo/model"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	InitMySQL()
	err = model.CreateTableUser()
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
	model.MySQL.Database = os.Getenv("DB_NAME")
	model.MySQL.Password = os.Getenv("DB_PASSWORD")
	model.MySQL.User = os.Getenv("DB_USER")
	model.MySQL.Host = os.Getenv("DB_HOST")
}
