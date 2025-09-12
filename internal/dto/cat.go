package dto

type UpdateCatParams struct {
	ID     string
	Salary float64
}

type CreateCatParams struct {
	Name              string
	YearsOfExperience int
	Breed             string
	Salary            float64
}
