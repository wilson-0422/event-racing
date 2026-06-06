package models

type Score struct {
	ID          int64  `json:"id"`
	AthleteID   int64  `json:"athlete_id"`
	ScheduleID  int64  `json:"schedule_id"`
	Score       string `json:"score"`
	Rank        int    `json:"rank"`
	Remark      string `json:"remark"`
	AthleteName string `json:"athlete_name"`
	CreatedAt   string `json:"created_at"`
}
