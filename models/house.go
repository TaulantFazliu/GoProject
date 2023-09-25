package models

type House struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PersonId int    `json:"person_id"`
}
