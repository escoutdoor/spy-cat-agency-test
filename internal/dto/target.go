package dto

type CreateTargetParams struct {
	Name    string
	Country string
}

type UpdateTargetParams struct {
	ID        string
	Completed *bool
	Notes     *string
}
