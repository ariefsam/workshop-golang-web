package model_test

import (
	"log"
	"testing"
	"todo/model"
)

func TestGetUserByEmail(t *testing.T) {
	InitMySQL()
	model.CreateTableUser()
	model.CreateUser("budi@gmail.com", "12345")
	var user *model.User
	user, err := model.GetUserByEmail("budi@gmail.com")
	if err != nil {
		t.Error(err.Error())
	}
	if user == nil {
		t.Error("User kosong")
	}
	log.Printf("%+v", user)
}
