package model

type Task struct {
	ID          int    `json:"ID"`
	Title       string `json:"title" required:"true"`
	Description string `json:"description" required:"true"`
	Status      string `json:"status"`
	Datecreated string `json:"datecreated"`
}
