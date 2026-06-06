package models

type Venue struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Capacity  int    `json:"capacity"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
