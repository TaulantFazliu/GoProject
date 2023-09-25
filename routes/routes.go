package routes

import (
	"GoProjects/routes/route"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.JSONLoggerMiddleware)

	// Define routes
	r.HandleFunc("/people", handlers.GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", handlers.GetPerson).Methods("GET")
	r.HandleFunc("/people", handlers.CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", handlers.UpdatePerson).Methods("PUT")
	r.HandleFunc("/people/{id}", handlers.DeletePerson).Methods("DELETE")

	r.HandleFunc("/houses", handlers.GetHouses).Methods("GET")

	return r
}
