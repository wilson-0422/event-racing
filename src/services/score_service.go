package services

import (
	"event-racing/src/config"
	"event-racing/src/models"
)

func GetAllScores() ([]models.Score, error) {
	rows, err := config.DB.Query(
		"SELECT s.id, s.athlete_id, s.schedule_id, s.score, s.rank, s.remark, a.name, s.created_at FROM scores s LEFT JOIN athletes a ON s.athlete_id = a.id ORDER BY s.id DESC",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var scores []models.Score
	for rows.Next() {
		var s models.Score
		if err := rows.Scan(&s.ID, &s.AthleteID, &s.ScheduleID, &s.Score, &s.Rank, &s.Remark, &s.AthleteName, &s.CreatedAt); err != nil {
			return nil, err
		}
		scores = append(scores, s)
	}
	return scores, nil
}

func GetScoreByID(id int64) (*models.Score, error) {
	var s models.Score
	err := config.DB.QueryRow(
		"SELECT s.id, s.athlete_id, s.schedule_id, s.score, s.rank, s.remark, a.name, s.created_at FROM scores s LEFT JOIN athletes a ON s.athlete_id = a.id WHERE s.id = ?",
		id,
	).Scan(&s.ID, &s.AthleteID, &s.ScheduleID, &s.Score, &s.Rank, &s.Remark, &s.AthleteName, &s.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func CreateScore(s *models.Score) error {
	result, err := config.DB.Exec(
		"INSERT INTO scores (athlete_id, schedule_id, score, rank, remark) VALUES (?, ?, ?, ?, ?)",
		s.AthleteID, s.ScheduleID, s.Score, s.Rank, s.Remark,
	)
	if err != nil {
		return err
	}
	s.ID, _ = result.LastInsertId()
	return nil
}

func UpdateScore(s *models.Score) error {
	_, err := config.DB.Exec(
		"UPDATE scores SET athlete_id=?, schedule_id=?, score=?, rank=?, remark=? WHERE id=?",
		s.AthleteID, s.ScheduleID, s.Score, s.Rank, s.Remark, s.ID,
	)
	return err
}

func DeleteScore(id int64) error {
	_, err := config.DB.Exec("DELETE FROM scores WHERE id = ?", id)
	return err
}

func GetScoresBySchedule(scheduleID int64) ([]models.Score, error) {
	rows, err := config.DB.Query(
		"SELECT s.id, s.athlete_id, s.schedule_id, s.score, s.rank, s.remark, a.name, s.created_at FROM scores s LEFT JOIN athletes a ON s.athlete_id = a.id WHERE s.schedule_id = ? ORDER BY s.rank",
		scheduleID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var scores []models.Score
	for rows.Next() {
		var s models.Score
		if err := rows.Scan(&s.ID, &s.AthleteID, &s.ScheduleID, &s.Score, &s.Rank, &s.Remark, &s.AthleteName, &s.CreatedAt); err != nil {
			return nil, err
		}
		scores = append(scores, s)
	}
	return scores, nil
}

func GetScoresByAthlete(athleteID int64) ([]models.Score, error) {
	rows, err := config.DB.Query(
		"SELECT s.id, s.athlete_id, s.schedule_id, s.score, s.rank, s.remark, a.name, s.created_at FROM scores s LEFT JOIN athletes a ON s.athlete_id = a.id WHERE s.athlete_id = ? ORDER BY s.id DESC",
		athleteID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var scores []models.Score
	for rows.Next() {
		var s models.Score
		if err := rows.Scan(&s.ID, &s.AthleteID, &s.ScheduleID, &s.Score, &s.Rank, &s.Remark, &s.AthleteName, &s.CreatedAt); err != nil {
			return nil, err
		}
		scores = append(scores, s)
	}
	return scores, nil
}
