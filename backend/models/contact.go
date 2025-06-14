package models

type Contact struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Relation string `json:"relation"`
}
