package api

import (
	"encoding/json"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {

	response := map[string]interface{}{
		"error": "",
	}

	var data struct {
		Token string `json:"token"`
	}
	json.NewDecoder(r.Body).Decode(&data)

	user, _ := GetSession(data.Token)
	if user.Email == "" {
		response["error"] = "No Auth"
		goto resp
	}
	response["user"] = map[string]interface{}{
		"email": user.Email,
	}

resp:

	dataResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataResponse)
	return
}
