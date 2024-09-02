package db

type Role struct {
	ID          string `json:"id"`
	Permissions int    `json:"permissions"`
	Name        string `json:"name"`
}
