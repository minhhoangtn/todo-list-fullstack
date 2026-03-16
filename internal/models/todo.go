package models

type Todo struct {
	ID        string `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}
