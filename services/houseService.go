package services

import (
	"GoProjects/models"
	"GoProjects/repositories"
)

type HouseService struct {
	HouseRepo *repositories.HouseRepository
}

func (hs *HouseService) GetAll() ([]models.House, error) {
	return hs.HouseRepo.GetAll()
}

func (hs *HouseService) GetById(id int) (*models.House, error) {
	return hs.HouseRepo.GetById(id)
}

func (hs *HouseService) Create(house *models.House) (*models.House, error) {
	return hs.HouseRepo.Create(house)
}

func (hs *HouseService) Update(house *models.House) error {
	return hs.HouseRepo.Update(house)
}

func (hs *HouseService) Delete(id int) error {
	return hs.HouseRepo.Delete(id)
}
