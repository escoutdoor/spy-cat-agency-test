package dto

type CreateMissionParams struct {
	Targets []CreateTargetParams
}

type UpdateMissionParams struct {
	ID        string
	CatID     *string
	Completed *bool
}
