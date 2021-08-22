package model

type Task struct {
	ID          int
	Name        string
	Description string
	Status      string
	OwnerID     int
}
