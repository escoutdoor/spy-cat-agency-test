package errors

import (
	"errors"
	"fmt"
)

var (
	TargetsLimitErr   = newError(TargetLimit, "there can only be 1-3 targets per mission")
	BreedDoesNotExist = newError(CatBreedDoesNotExist, "there is no such cat breed")
	NoFieldsToUpdate  = newError(NoFieldsNeedToBeUpdated, "no fields to update")
)

type Error struct {
	Code Code
	Err  error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func newError(code Code, err string) *Error {
	return &Error{
		Code: code,
		Err:  errors.New(err),
	}
}

func CatNotFoundWithID(catID string) *Error {
	msg := fmt.Sprintf("no cat found with id %q", catID)

	return newError(CatNotFound, msg)
}

func MissionNotFoundWithID(missionID string) *Error {
	msg := fmt.Sprintf("no mission found with id %q", missionID)

	return newError(MissionNotFound, msg)
}

func MissionCannotBeDeleted(missionID string) *Error {
	msg := fmt.Sprintf("mission with id %q cannot be deleted because it's already assigned to a cat", missionID)

	return newError(MissionCannotBeDeletedAssignedToCat, msg)
}

func TargetNotFoundWithID(targetID string) *Error {
	msg := fmt.Sprintf("no target found with id %q", targetID)

	return newError(TargetNotFound, msg)
}

func MissionAlreadyCompletedWithID(missionID string) *Error {
	msg := fmt.Sprintf("mission with id %q is already completed", missionID)

	return newError(MissionAlreadyCompleted, msg)
}

func CatOnMissionWithID(catID string) *Error {
	msg := fmt.Sprintf("cat with id %q is already on a mission", catID)

	return newError(CatOnMission, msg)
}

func TargetAlreadyCompletedWithID(targetID string) *Error {
	msg := fmt.Sprintf("target with id %q is already completed", targetID)

	return newError(TargetAlreadyCompleted, msg)
}
