package services

import (
	"GoProjects/models"
)

type HouseAssociationService struct{}

func NewHouseAssociationService() *HouseAssociationService {
	return &HouseAssociationService{}
}

func (has *HouseAssociationService) AssociateHousesWithPeople(people []*models.Person, houses []*models.House) {
	personHousesMap := make(map[int][]*models.House)

	for _, house := range houses {
		personID := house.PersonId
		if _, exists := personHousesMap[personID]; !exists {
			personHousesMap[personID] = []*models.House{}
		}
		personHousesMap[personID] = append(personHousesMap[personID], house)
	}

	for _, person := range people {
		if houses, exists := personHousesMap[person.ID]; exists {
			person.HouseCount = len(houses)
		} else {
			person.HouseCount = 0
		}
	}
}
