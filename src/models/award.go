package models

type Award struct {
	ID              int64  `json:"id"`
	AthleteID       int64  `json:"athlete_id"`
	Event           string `json:"event"`
	MedalType       string `json:"medal_type"`
	CompetitionName string `json:"competition_name"`
	AthleteName     string `json:"athlete_name"`
	CreatedAt       string `json:"created_at"`
}
