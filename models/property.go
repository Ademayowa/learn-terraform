package models

import (
	"github.com/Ademayowa/learn-terraform/db"
	"github.com/google/uuid"
)

type Property struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Location string `json:"location"`
}

// Save property into the database
func (p *Property) Save() error {
	p.ID = uuid.New().String()

	query := `INSERT INTO properties(id, title, location) VALUES($1, $2, $3)`

	_, err := db.DB.Exec(query, p.ID, p.Title, p.Location)
	return err
}

// Retrieves all properties from database
func GetAllProperties() ([]Property, error) {
	query := `SELECT id, title, location FROM properties`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var properties []Property
	for rows.Next() {
		var p Property
		err := rows.Scan(&p.ID, &p.Title, &p.Location)
		if err != nil {
			return nil, err
		}
		properties = append(properties, p)
	}

	return properties, rows.Err()
}
