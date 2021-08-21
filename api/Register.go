package api

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/model"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&data)

	id, err := model.CreateUser(data.Email, data.Password)

	//log.Printf("%+v", data)
	log.Println(id, err)
}
