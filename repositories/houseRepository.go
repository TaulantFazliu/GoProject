package repositories

import (
	"GoProjects/models"
	"database/sql"
)

//// TableBRepository defines the methods for interacting with TableB.
//type TableBRepository interface {
//	GetByID(id string) (*models.TableB, error)
//	Create(tableB *models.TableB) error
//	Update(tableB *models.TableB) error
//	Delete(id string) error
//}

type HouseRepository struct {
	db *sql.DB // Or use a database connection pool if preferred
}

func (hr *HouseRepository) GetAll() ([]models.House, error) {
	var houses []models.House

	rows, err := hr.db.Query("SELECT * FROM houses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var h models.House
		if err := rows.Scan(&h.ID, &h.Name); err != nil {
			return nil, err
		}
		houses = append(houses, h)
	}

	return houses, nil
}

func (hr *HouseRepository) Create(house *models.House) (*models.House, error) {
	err := hr.db.QueryRow("INSERT INTO houses (name) VALUES ($1) RETURNING id",
		house.Name).Scan(&house.ID)

	if err != nil {
		return nil, err
	}

	return house, nil
}

func (hr *HouseRepository) GetById(id int) (*models.House, error) {
	var house models.House

	err := hr.db.QueryRow("SELECT * FROM houses WHERE id = $1", id).
		Scan(&house.ID, &house.Name)

	if err != nil {
		return nil, err
	}

	return &house, nil
}

func (hr *HouseRepository) Update(house *models.House) error {
	_, err := hr.db.Exec("UPDATE houses SET name = $1 WHERE id = $2",
		house.Name, house.ID)

	return err
}

func (hr *HouseRepository) Delete(id int) error {
	_, err := hr.db.Exec("DELETE FROM houses WHERE id = $1", id)
	return err
}
