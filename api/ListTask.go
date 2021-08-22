package api

import (
	"encoding/json"
	"net/http"
	"todo/model"
)

func ListTask(w http.ResponseWriter, r *http.Request) {
	var tasks []model.Task
	response := map[string]interface{}{
		"error": "",
	}
	var data struct {
		Token string `json:"token"`
	}
	json.NewDecoder(r.Body).Decode(&data)

	user, err := GetSession(data.Token)
	if err != nil {
		response["error"] = err
		goto resp
	}
	if user.ID == 0 {
		response["error"] = "Unauth"
		goto resp
	}

	tasks, err = model.ListTask(user.ID)
	if err != nil {
		response["error"] = err
		goto resp
	}
	response["data"] = tasks
resp:

	dataResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataResponse)
	return

}
