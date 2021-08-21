package model_test

import (
	"log"
	"testing"
	"todo/model"
)

func TestGetUID(t *testing.T) {
	str := model.GetUID()
	if str == "" {
		t.Error("uid kosong")
	}
	log.Println(str)
}
