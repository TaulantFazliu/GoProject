package repositories

import (
	"GoProjects/models"
	"database/sql"
)

type PersonRepository struct {
	db *sql.DB // Or use a database connection pool if preferred
}

func (pr *PersonRepository) GetAll() ([]models.Person, error) {
	var people []models.Person

	rows, err := pr.db.Query("SELECT * FROM people")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Person
		if err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.HouseID); err != nil {
			return nil, err
		}
		people = append(people, p)
	}

	return people, nil
}

func (pr *PersonRepository) GetById(id int) (*models.Person, error) {
	var person models.Person

	err := pr.db.QueryRow("SELECT * FROM people WHERE id = $1", id).
		Scan(&person.ID, &person.Name, &person.Age, &person.HouseID)

	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (pr *PersonRepository) Update(person *models.Person) error {
	_, err := pr.db.Exec("UPDATE people SET name = $1, age = $2, house_id = $3 WHERE id = $4",
		person.Name, person.Age, person.HouseID, person.ID)

	return err
}

func (pr *PersonRepository) Delete(id int) error {
	_, err := pr.db.Exec("DELETE FROM people WHERE id = $1", id)
	return err
}

func (pr *PersonRepository) Create(person *models.Person) (*models.Person, error) {
	err := pr.db.QueryRow("INSERT INTO people (name, age, house_id) VALUES ($1, $2, $3) RETURNING id",
		person.Name, person.Age, person.HouseID).Scan(&person.ID)

	if err != nil {
		return nil, err
	}

	return person, nil
}

//// NewUserRepository creates a new instance of UserRepository.
//func NewUserRepository(db *sql.DB) *PersonRepository {
//	return &PersonRepository{db: db}
//}
//
//// GetUserByID retrieves a user by their ID using a raw SQL query.
//func (r *PersonRepository) GetUserByID(id int) (*models.Person, error) {
//	// Define the SQL query
//	query := "SELECT id, name, age FROM users WHERE id = $1"
//
//	// Execute the SQL query
//	row := r.db.QueryRow(query, id)
//
//	// Parse the query result into a User struct
//	var person models.Person
//	if err := row.Scan(&person.ID, &person.Name, &person.Age); err != nil {
//		return nil, err
//	}
//
//	return &person, nil
//}
