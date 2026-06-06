package services

import (
	"event-racing/src/config"
	"event-racing/src/models"
)

func GetAllAthletes() ([]models.Athlete, error) {
	rows, err := config.DB.Query("SELECT id, name, gender, age, team, event, phone, id_number, created_at FROM athletes ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var athletes []models.Athlete
	for rows.Next() {
		var a models.Athlete
		if err := rows.Scan(&a.ID, &a.Name, &a.Gender, &a.Age, &a.Team, &a.Event, &a.Phone, &a.IDNumber, &a.CreatedAt); err != nil {
			return nil, err
		}
		athletes = append(athletes, a)
	}
	return athletes, nil
}

func GetAthleteByID(id int64) (*models.Athlete, error) {
	var a models.Athlete
	err := config.DB.QueryRow(
		"SELECT id, name, gender, age, team, event, phone, id_number, created_at FROM athletes WHERE id = ?",
		id,
	).Scan(&a.ID, &a.Name, &a.Gender, &a.Age, &a.Team, &a.Event, &a.Phone, &a.IDNumber, &a.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func CreateAthlete(a *models.Athlete) error {
	result, err := config.DB.Exec(
		"INSERT INTO athletes (name, gender, age, team, event, phone, id_number) VALUES (?, ?, ?, ?, ?, ?, ?)",
		a.Name, a.Gender, a.Age, a.Team, a.Event, a.Phone, a.IDNumber,
	)
	if err != nil {
		return err
	}
	a.ID, _ = result.LastInsertId()
	return nil
}

func UpdateAthlete(a *models.Athlete) error {
	_, err := config.DB.Exec(
		"UPDATE athletes SET name=?, gender=?, age=?, team=?, event=?, phone=?, id_number=? WHERE id=?",
		a.Name, a.Gender, a.Age, a.Team, a.Event, a.Phone, a.IDNumber, a.ID,
	)
	return err
}

func DeleteAthlete(id int64) error {
	_, err := config.DB.Exec("DELETE FROM athletes WHERE id = ?", id)
	return err
}

func GetAthletesByEvent(event string) ([]models.Athlete, error) {
	rows, err := config.DB.Query(
		"SELECT id, name, gender, age, team, event, phone, id_number, created_at FROM athletes WHERE event = ? ORDER BY name",
		event,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var athletes []models.Athlete
	for rows.Next() {
		var a models.Athlete
		if err := rows.Scan(&a.ID, &a.Name, &a.Gender, &a.Age, &a.Team, &a.Event, &a.Phone, &a.IDNumber, &a.CreatedAt); err != nil {
			return nil, err
		}
		athletes = append(athletes, a)
	}
	return athletes, nil
}

func CountAthletes() (int, error) {
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM athletes").Scan(&count)
	return count, err
}
