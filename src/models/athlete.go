package models

type Athlete struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Team      string `json:"team"`
	Event     string `json:"event"`
	Phone     string `json:"phone"`
	IDNumber  string `json:"id_number"`
	CreatedAt string `json:"created_at"`
}
