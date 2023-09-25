package handlers

import (
	"GoProjects/services"
	"github.com/gorilla/mux"
)

type AppHandlers struct {
	PersonHandlers      *PersonHandlers
	HouseHandlers       *HouseHandlers
	AssociationHandlers *AssociationHandlers
}

func NewAppHandlers(personService *services.PersonService, houseService *services.HouseService, associationService *services.HouseAssociationService) *AppHandlers {
	return &AppHandlers{
		PersonHandlers:      NewPersonHandlers(personService),
		HouseHandlers:       NewHouseHandlers(houseService),
		AssociationHandlers: NewAssociationHandlers(associationService),
	}
}

func (ah *AppHandlers) RegisterRoutes(router *mux.Router) {
	// Person
	router.HandleFunc("/people", ah.PersonHandlers.GetAllPeople).Methods("GET")
	router.HandleFunc("/people/{id:[0-9]+}", ah.PersonHandlers.GetOnePersonById).Methods("GET")
	router.HandleFunc("/people", ah.PersonHandlers.CreateOnePerson).Methods("POST")
	router.HandleFunc("/people/{id:[0-9]+}", ah.PersonHandlers.UpdateOnePerson).Methods("PUT")
	router.HandleFunc("/people/{id:[0-9]+}", ah.PersonHandlers.DeleteOnePerson).Methods("DELETE")

	// House
	router.HandleFunc("/houses", ah.HouseHandlers.GetAllHouses).Methods("GET")
	router.HandleFunc("/houses/{id:[0-9]+}", ah.HouseHandlers.GetOneHouseById).Methods("GET")
	router.HandleFunc("/houses", ah.HouseHandlers.CreateOneHouse).Methods("POST")
	router.HandleFunc("/houses/{id:[0-9]+}", ah.HouseHandlers.UpdateOneHouse).Methods("PUT")
	router.HandleFunc("/houses/{id:[0-9]+}", ah.HouseHandlers.DeleteOneHouse).Methods("DELETE")

	//Association
	router.HandleFunc("/associate-houses", ah.AssociationHandlers.AssociateHousesWithPeopleHandler).Methods("POST")

}
