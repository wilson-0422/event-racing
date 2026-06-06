package services

import (
	"event-racing/src/config"
	"event-racing/src/models"
)

func GetAllAwards() ([]models.Award, error) {
	rows, err := config.DB.Query(
		"SELECT a.id, a.athlete_id, a.event, a.medal_type, a.competition_name, at.name, a.created_at FROM awards a LEFT JOIN athletes at ON a.athlete_id = at.id ORDER BY a.id DESC",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var awards []models.Award
	for rows.Next() {
		var a models.Award
		if err := rows.Scan(&a.ID, &a.AthleteID, &a.Event, &a.MedalType, &a.CompetitionName, &a.AthleteName, &a.CreatedAt); err != nil {
			return nil, err
		}
		awards = append(awards, a)
	}
	return awards, nil
}

func GetAwardByID(id int64) (*models.Award, error) {
	var a models.Award
	err := config.DB.QueryRow(
		"SELECT a.id, a.athlete_id, a.event, a.medal_type, a.competition_name, at.name, a.created_at FROM awards a LEFT JOIN athletes at ON a.athlete_id = at.id WHERE a.id = ?",
		id,
	).Scan(&a.ID, &a.AthleteID, &a.Event, &a.MedalType, &a.CompetitionName, &a.AthleteName, &a.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func CreateAward(a *models.Award) error {
	result, err := config.DB.Exec(
		"INSERT INTO awards (athlete_id, event, medal_type, competition_name) VALUES (?, ?, ?, ?)",
		a.AthleteID, a.Event, a.MedalType, a.CompetitionName,
	)
	if err != nil {
		return err
	}
	a.ID, _ = result.LastInsertId()
	return nil
}

func DeleteAward(id int64) error {
	_, err := config.DB.Exec("DELETE FROM awards WHERE id = ?", id)
	return err
}

func GetAwardsByAthlete(athleteID int64) ([]models.Award, error) {
	rows, err := config.DB.Query(
		"SELECT a.id, a.athlete_id, a.event, a.medal_type, a.competition_name, at.name, a.created_at FROM awards a LEFT JOIN athletes at ON a.athlete_id = at.id WHERE a.athlete_id = ? ORDER BY a.id DESC",
		athleteID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var awards []models.Award
	for rows.Next() {
		var a models.Award
		if err := rows.Scan(&a.ID, &a.AthleteID, &a.Event, &a.MedalType, &a.CompetitionName, &a.AthleteName, &a.CreatedAt); err != nil {
			return nil, err
		}
		awards = append(awards, a)
	}
	return awards, nil
}

func CountAwards() (int, error) {
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM awards").Scan(&count)
	return count, err
}
