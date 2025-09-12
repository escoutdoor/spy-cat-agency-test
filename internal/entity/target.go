package entity

import "time"

type Target struct {
	ID        string
	MissionID string
	Name      string
	Country   string
	Notes     string
	Completed bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
