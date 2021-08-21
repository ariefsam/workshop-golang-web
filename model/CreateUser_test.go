package model_test

import (
	"log"
	"testing"
	"todo/model"
)

func TestCreateUser(t *testing.T) {
	InitMySQL()
	id, err := model.CreateUser("arief@gmail.com", "12345")
	log.Println(id)
	if err != nil {
		t.Error(err.Error())
	}
}
