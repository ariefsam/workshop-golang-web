package api

import (
	"encoding/json"
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
	response := map[string]interface{}{
		"id":    id,
		"error": "",
	}
	if err != nil {
		response["error"] = err.Error()
	}

	dataResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataResponse)
	return
}
