package main

import (
	"log"
	"net/http"
	"time"
	"todo/api"
	"todo/page"
)

func Server() {
	http.HandleFunc("/api/register", api.Register)
	http.HandleFunc("/api/login", api.Login)
	http.HandleFunc("/api/profile", api.Profile)
	http.HandleFunc("/api/create-task", api.CreateTask)
	http.HandleFunc("/api/list-task", api.ListTask)
	http.HandleFunc("/api/edit-task", api.EditTask)
	http.HandleFunc("/api/change-task-status", api.ChangeTaskStatus)
	http.HandleFunc("/register", page.Register)
	http.HandleFunc("/dashboard", page.Dashboard)
	http.HandleFunc("/login", page.Login)
	http.HandleFunc("/", page.Index)
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
