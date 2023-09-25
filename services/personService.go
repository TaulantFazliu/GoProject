package services

import (
	"GoProjects/models"
	"GoProjects/repositories"
)

type PersonService struct {
	PersonRepo *repositories.PersonRepository
}

func (ps *PersonService) GetAll() ([]models.Person, error) {
	return ps.PersonRepo.GetAll()
}

func (ps *PersonService) GetById(id int) (*models.Person, error) {
	return ps.PersonRepo.GetById(id)
}

func (ps *PersonService) Create(person *models.Person) (*models.Person, error) {
	return ps.PersonRepo.Create(person)
}

func (ps *PersonService) Update(person *models.Person) error {
	return ps.PersonRepo.Update(person)
}

func (ps *PersonService) Delete(id int) error {
	return ps.PersonRepo.Delete(id)
}
