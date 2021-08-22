package model_test

import (
	"log"
	"testing"
	"todo/model"
)

func TestCreateTask(t *testing.T) {
	InitMySQL()
	var task model.Task
	task.ID = 0
	task.Name = "Task Pertama"
	task.Description = "Task yang dibuat pertama kali"
	task.Status = "todo"
	task.OwnerID = 1

	taskID, err := model.CreateTask(task)
	if err != nil {
		t.Error(err.Error())
	}
	if taskID == 0 {
		t.Error("Task ID 0")
	}
	log.Println(taskID)
}
