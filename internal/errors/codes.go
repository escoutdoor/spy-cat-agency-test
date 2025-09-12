package errors

type Code string

const (
	CatNotFound                         Code = "CAT_NOT_FOUND"
	CatBreedDoesNotExist                Code = "CAT_BREED_DOES_NOT_EXIST"
	MissionNotFound                     Code = "MISSION_NOT_FOUND"
	CatOnMission                        Code = "CAT_ON_MISSION"
	TargetNotFound                      Code = "TARGET_NOT_FOUND"
	MissionCannotBeDeletedAssignedToCat Code = "MISSION_CANNOT_DELETED_ASSIGNED_TO_CAT"
	NoFieldsNeedToBeUpdated             Code = "NO_FIELDS_NEED_TO_BE_UPDATED"
	MissionAlreadyCompleted             Code = "MISSION_ALREADY_COMPLETED"
	TargetAlreadyCompleted              Code = "TARGET_ALREADY_COMPLETED"
	TargetLimit                         Code = "TARGET_LIMIT"
)
