package services

import (
	"event-racing/src/config"
	"event-racing/src/models"
)

func GetAllSchedules() ([]models.Schedule, error) {
	rows, err := config.DB.Query(
		"SELECT s.id, s.group_id, s.venue_id, g.name, v.name, s.start_time, s.end_time, s.status, s.created_at FROM schedules s LEFT JOIN groups g ON s.group_id = g.id LEFT JOIN venues v ON s.venue_id = v.id ORDER BY s.start_time DESC",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var schedules []models.Schedule
	for rows.Next() {
		var s models.Schedule
		if err := rows.Scan(&s.ID, &s.GroupID, &s.VenueID, &s.GroupName, &s.VenueName, &s.StartTime, &s.EndTime, &s.Status, &s.CreatedAt); err != nil {
			return nil, err
		}
		schedules = append(schedules, s)
	}
	return schedules, nil
}

func GetScheduleByID(id int64) (*models.Schedule, error) {
	var s models.Schedule
	err := config.DB.QueryRow(
		"SELECT s.id, s.group_id, s.venue_id, g.name, v.name, s.start_time, s.end_time, s.status, s.created_at FROM schedules s LEFT JOIN groups g ON s.group_id = g.id LEFT JOIN venues v ON s.venue_id = v.id WHERE s.id = ?",
		id,
	).Scan(&s.ID, &s.GroupID, &s.VenueID, &s.GroupName, &s.VenueName, &s.StartTime, &s.EndTime, &s.Status, &s.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func CreateSchedule(s *models.Schedule) error {
	result, err := config.DB.Exec(
		"INSERT INTO schedules (group_id, venue_id, start_time, end_time, status) VALUES (?, ?, ?, ?, ?)",
		s.GroupID, s.VenueID, s.StartTime, s.EndTime, s.Status,
	)
	if err != nil {
		return err
	}
	s.ID, _ = result.LastInsertId()
	return nil
}

func UpdateSchedule(s *models.Schedule) error {
	_, err := config.DB.Exec(
		"UPDATE schedules SET group_id=?, venue_id=?, start_time=?, end_time=?, status=? WHERE id=?",
		s.GroupID, s.VenueID, s.StartTime, s.EndTime, s.Status, s.ID,
	)
	return err
}

func DeleteSchedule(id int64) error {
	_, err := config.DB.Exec("DELETE FROM schedules WHERE id = ?", id)
	return err
}

func GetSchedulesByVenue(venueID int64) ([]models.Schedule, error) {
	rows, err := config.DB.Query(
		"SELECT s.id, s.group_id, s.venue_id, g.name, v.name, s.start_time, s.end_time, s.status, s.created_at FROM schedules s LEFT JOIN groups g ON s.group_id = g.id LEFT JOIN venues v ON s.venue_id = v.id WHERE s.venue_id = ? ORDER BY s.start_time",
		venueID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var schedules []models.Schedule
	for rows.Next() {
		var s models.Schedule
		if err := rows.Scan(&s.ID, &s.GroupID, &s.VenueID, &s.GroupName, &s.VenueName, &s.StartTime, &s.EndTime, &s.Status, &s.CreatedAt); err != nil {
			return nil, err
		}
		schedules = append(schedules, s)
	}
	return schedules, nil
}

func CountSchedules() (int, error) {
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM schedules").Scan(&count)
	return count, err
}
