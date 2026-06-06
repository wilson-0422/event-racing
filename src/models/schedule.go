package models

type Schedule struct {
	ID        int64  `json:"id"`
	GroupID   int64  `json:"group_id"`
	VenueID   int64  `json:"venue_id"`
	GroupName string `json:"group_name"`
	VenueName string `json:"venue_name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
