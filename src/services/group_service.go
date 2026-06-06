package services

import (
	"event-racing/src/config"
	"event-racing/src/models"
)

func GetAllGroups() ([]models.Group, error) {
	rows, err := config.DB.Query("SELECT id, name, event, status, created_at FROM groups ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var groups []models.Group
	for rows.Next() {
		var g models.Group
		if err := rows.Scan(&g.ID, &g.Name, &g.Event, &g.Status, &g.CreatedAt); err != nil {
			return nil, err
		}
		athletes, _ := GetAthletesByGroupID(g.ID)
		g.Athletes = athletes
		groups = append(groups, g)
	}
	return groups, nil
}

func GetGroupByID(id int64) (*models.Group, error) {
	var g models.Group
	err := config.DB.QueryRow(
		"SELECT id, name, event, status, created_at FROM groups WHERE id = ?",
		id,
	).Scan(&g.ID, &g.Name, &g.Event, &g.Status, &g.CreatedAt)
	if err != nil {
		return nil, err
	}
	athletes, _ := GetAthletesByGroupID(g.ID)
	g.Athletes = athletes
	return &g, nil
}

func GetAthletesByGroupID(groupID int64) ([]models.Athlete, error) {
	rows, err := config.DB.Query(
		"SELECT a.id, a.name, a.gender, a.age, a.team, a.event, a.phone, a.id_number, a.created_at FROM athletes a INNER JOIN group_athletes ga ON a.id = ga.athlete_id WHERE ga.group_id = ? ORDER BY a.name",
		groupID,
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

func CreateGroup(g *models.Group) error {
	result, err := config.DB.Exec(
		"INSERT INTO groups (name, event, status) VALUES (?, ?, ?)",
		g.Name, g.Event, g.Status,
	)
	if err != nil {
		return err
	}
	g.ID, _ = result.LastInsertId()
	return nil
}

func UpdateGroup(g *models.Group) error {
	_, err := config.DB.Exec(
		"UPDATE groups SET name=?, event=?, status=? WHERE id=?",
		g.Name, g.Event, g.Status, g.ID,
	)
	return err
}

func DeleteGroup(id int64) error {
	_, err := config.DB.Exec("DELETE FROM group_athletes WHERE group_id = ?", id)
	if err != nil {
		return err
	}
	_, err = config.DB.Exec("DELETE FROM groups WHERE id = ?", id)
	return err
}

func ArrangeAthletes(groupID int64, athleteIDs []int64) error {
	_, err := config.DB.Exec("DELETE FROM group_athletes WHERE group_id = ?", groupID)
	if err != nil {
		return err
	}
	for _, aid := range athleteIDs {
		if _, err := config.DB.Exec("INSERT INTO group_athletes (group_id, athlete_id) VALUES (?, ?)", groupID, aid); err != nil {
			return err
		}
	}
	_, err = config.DB.Exec("UPDATE groups SET status = 'arranged' WHERE id = ?", groupID)
	return err
}

func CountGroups() (int, error) {
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM groups").Scan(&count)
	return count, err
}
