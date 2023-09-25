package handlers

import (
	"GoProjects/models"
	"GoProjects/services"
	"encoding/json"
	"net/http"
)

type AssociationHandlers struct {
	HouseAssociationService *services.HouseAssociationService
}

func NewAssociationHandlers(associationService *services.HouseAssociationService) *AssociationHandlers {
	return &AssociationHandlers{HouseAssociationService: associationService}
}

func (ah *AssociationHandlers) AssociateHousesWithPeopleHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch people and houses from your database or other sources
	var people []*models.Person
	var houses []*models.House

	// Create an instance of HouseAssociationService
	houseAssocService := services.NewHouseAssociationService()

	// Call the service method to associate houses with people
	houseAssocService.AssociateHousesWithPeople(people, houses)

	// Respond with a success message or appropriate JSON response
	response := map[string]string{"message": "Houses associated with people and counts updated successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
