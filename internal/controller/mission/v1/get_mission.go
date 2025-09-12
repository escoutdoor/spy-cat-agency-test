package v1

import (
	"errors"

	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
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
		appErr := new(apperrors.Error)
		if errors.As(err, &appErr) {
			if appErr.Code == apperrors.MissionNotFound {
				return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": appErr.Error(),
				})
			}
		}

		return err
	}

	type response struct {
		Mission missionResponse `json:"mission"`
	}
	return ctx.Status(fiber.StatusOK).JSON(response{Mission: missionToMissionResponse(mission)})
}
