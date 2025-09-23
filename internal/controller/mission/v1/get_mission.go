package v1

import (
	"github.com/gofiber/fiber/v3"
)

func (c *controller) getMission(ctx fiber.Ctx) error {
	missionID := ctx.Params(missionIDParam)
	if err := validateMissionID(missionID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	mission, err := c.missionService.GetMission(ctx, missionID)
	if err != nil {
		return err
	}

	type response struct {
		Mission missionResponse `json:"mission"`
	}
	return ctx.Status(fiber.StatusOK).JSON(response{Mission: missionToMissionResponse(mission)})
}
