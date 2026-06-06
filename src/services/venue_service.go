package services

import (
	"event-racing/src/config"
	"event-racing/src/models"
)

func GetAllVenues() ([]models.Venue, error) {
	rows, err := config.DB.Query("SELECT id, name, location, capacity, status, created_at FROM venues ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var venues []models.Venue
	for rows.Next() {
		var v models.Venue
		if err := rows.Scan(&v.ID, &v.Name, &v.Location, &v.Capacity, &v.Status, &v.CreatedAt); err != nil {
			return nil, err
		}
		venues = append(venues, v)
	}
	return venues, nil
}

func GetVenueByID(id int64) (*models.Venue, error) {
	var v models.Venue
	err := config.DB.QueryRow(
		"SELECT id, name, location, capacity, status, created_at FROM venues WHERE id = ?",
		id,
	).Scan(&v.ID, &v.Name, &v.Location, &v.Capacity, &v.Status, &v.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func CreateVenue(v *models.Venue) error {
	result, err := config.DB.Exec(
		"INSERT INTO venues (name, location, capacity, status) VALUES (?, ?, ?, ?)",
		v.Name, v.Location, v.Capacity, v.Status,
	)
	if err != nil {
		return err
	}
	v.ID, _ = result.LastInsertId()
	return nil
}

func UpdateVenue(v *models.Venue) error {
	_, err := config.DB.Exec(
		"UPDATE venues SET name=?, location=?, capacity=?, status=? WHERE id=?",
		v.Name, v.Location, v.Capacity, v.Status, v.ID,
	)
	return err
}

func DeleteVenue(id int64) error {
	_, err := config.DB.Exec("DELETE FROM venues WHERE id = ?", id)
	return err
}

func CountVenues() (int, error) {
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM venues").Scan(&count)
	return count, err
}
