package models

type Person struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	HouseCount int    `json:"house_count"`
}

//
//import (
//	"database/sql"
//	"encoding/json"
//	"log"
//	"net/http"
//)
//
//type Person struct {
//	ID     uint     `json:"id"`
//	Name   string   `json:"name"`
//	Age    int      `json:"age"`
//	Houses []TableB // One-to-Many relationship with TableB
//}
//
//func create(w http.ResponseWriter, r *http.Request) {
//	var body Person
//	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		log.Print("error decoding request body into Person struct %v", err)
//		return
//	}
//	if err := sql.DB.QueryRow("INSERT INTO persons(id,name, age)values (%1, %2, %3)", body.ID, body.Name, body.Age) {
//	}
//}
//
//func getAll(w http.ResponseWriter, _ *http.Request) {
//
//	var persons []Person
//
//	j, err := json.Marshal(persons)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		log.Printf("error marshalling persons into json %v", err)
//		return
//	}
//	w.Write(j)
//
//}
