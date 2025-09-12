package entity

import "time"

type Cat struct {
	ID                string
	Name              string
	YearsOfExperience int
	Breed             string
	Salary            float64

	CreatedAt time.Time
	UpdatedAt time.Time
}
