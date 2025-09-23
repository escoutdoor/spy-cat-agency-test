package errors

import (
	"errors"
	"fmt"

	"github.com/escoutdoor/spy-cat-agency-test/internal/errors/code"
)

var (
	TargetsLimitErr   = newError(code.TargetLimit, "there can only be 1-3 targets per mission")
	BreedDoesNotExist = newError(code.CatBreedDoesNotExist, "there is no such cat breed")
	NoFieldsToUpdate  = newError(code.NoFieldsNeedToBeUpdated, "no fields to update")
)

type Error struct {
	Code code.Code
	Err  error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func newError(code code.Code, err string) *Error {
	return &Error{
		Code: code,
		Err:  errors.New(err),
	}
}

func CatNotFoundWithID(catID string) *Error {
	msg := fmt.Sprintf("no cat found with id %q", catID)

	return newError(code.CatNotFound, msg)
}

func MissionNotFoundWithID(missionID string) *Error {
	msg := fmt.Sprintf("no mission found with id %q", missionID)

	return newError(code.MissionNotFound, msg)
}

func MissionCannotBeDeleted(missionID string) *Error {
	msg := fmt.Sprintf("mission with id %q cannot be deleted because it's already assigned to a cat", missionID)

	return newError(code.MissionCannotBeDeletedAssignedToCat, msg)
}

func TargetNotFoundWithID(targetID string) *Error {
	msg := fmt.Sprintf("no target found with id %q", targetID)

	return newError(code.TargetNotFound, msg)
}

func MissionAlreadyCompletedWithID(missionID string) *Error {
	msg := fmt.Sprintf("mission with id %q is already completed", missionID)

	return newError(code.MissionAlreadyCompleted, msg)
}

func CatOnMissionWithID(catID string) *Error {
	msg := fmt.Sprintf("cat with id %q is already on a mission", catID)

	return newError(code.CatOnMission, msg)
}

func TargetAlreadyCompletedWithID(targetID string) *Error {
	msg := fmt.Sprintf("target with id %q is already completed", targetID)

	return newError(code.TargetAlreadyCompleted, msg)
}
