package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"todo/model"
)

var Session map[string]model.User = make(map[string]model.User)
var m sync.Mutex

func Login(w http.ResponseWriter, r *http.Request) {
	var uid string
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
		goto resp
	}

	if data.Password != user.Password {
		response["error"] = "Wrong Password"
		goto resp
	}

	if err != nil {
		response["error"] = err.Error()
	}

	uid = model.GetUID()
	m.Lock()
	Session[uid] = *user
	log.Println(Session)
	m.Unlock()
	response["session_id"] = uid

resp:

	dataResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataResponse)
	return
}
