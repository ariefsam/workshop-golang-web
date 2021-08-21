package api

import (
	"encoding/json"
	"net/http"
	"todo/model"
)

func Login(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"error": "",
	}

	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&data)

	user, err := model.GetUserByEmail(data.Email)
	if user == nil {
		response["error"] = "User not found"
	}

	if data.Password != user.Password {
		response["error"] = "Wrong Password"
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
