package models

type Group struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Event     string    `json:"event"`
	Status    string    `json:"status"`
	Athletes  []Athlete `json:"athletes"`
	CreatedAt string    `json:"created_at"`
}

type GroupAthlete struct {
	GroupID   int64 `json:"group_id"`
	AthleteID int64 `json:"athlete_id"`
}
