package v1

import (
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
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
