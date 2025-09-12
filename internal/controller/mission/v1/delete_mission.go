package v1

import (
	"errors"

	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
	"github.com/gofiber/fiber/v3"
)

func (c *controller) deleteMission(ctx fiber.Ctx) error {
	missionID := ctx.Params(missionIDParam)
	if err := validateMissionID(missionID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.missionService.DeleteMission(ctx, missionID); err != nil {
		appErr := new(apperrors.Error)
		if errors.As(err, &appErr) {
			switch appErr.Code {
			case apperrors.MissionNotFound:
				return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": appErr.Error(),
				})
			case apperrors.MissionCannotBeDeletedAssignedToCat:
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": appErr.Error(),
				})
			}
		}

		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
