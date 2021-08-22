package api

import (
	"encoding/json"
	"net/http"
	"todo/model"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	var taskID int
	response := map[string]interface{}{
		"error": "",
	}

	var data struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      string `json:"status"`
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

	task.Name = data.Name
	task.Description = data.Description
	task.Status = data.Status
	task.OwnerID = user.ID

	taskID, err = model.CreateTask(task)
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
