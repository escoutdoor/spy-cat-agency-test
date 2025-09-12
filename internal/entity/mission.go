package entity

import "time"

type Mission struct {
	ID        string
	CatID     *string
	Targets   []Target
	Completed bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
