package api

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/model"
)

func EditTask(w http.ResponseWriter, r *http.Request) {
	var taskID int
	response := map[string]interface{}{
		"error": "",
	}

	var data struct {
		TaskID      int    `json:"task_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Token       string `json:"token"`
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
	log.Println(data)
	err = model.EditTask(data.TaskID, user.ID, data.Name, data.Description)
	if err != nil {
		response["error"] = err
		goto resp
	}
	response["taskID"] = taskID

resp:

	dataResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataResponse)
	return
}
