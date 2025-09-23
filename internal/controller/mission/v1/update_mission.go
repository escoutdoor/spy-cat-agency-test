package v1

import (
	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type updateMissionRequestBody struct {
	CatID *string `json:"catId" validate:"omitempty,uuid"`
}

func (c *controller) updateMission(ctx fiber.Ctx) error {
	missionID := ctx.Params(missionIDParam)
	if err := validateMissionID(missionID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var req updateMissionRequestBody
	if err := ctx.Bind().Body(&req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"field": e.Field(),
					"error": e.Error(),
				})
			}
		}

		return err
	}

	if err := c.missionService.UpdateMission(ctx, updateMissionRequestBodyToUpdateMissionParams(req, missionID)); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func updateMissionRequestBodyToUpdateMissionParams(req updateMissionRequestBody, missionID string) dto.UpdateMissionParams {
	return dto.UpdateMissionParams{
		ID:    missionID,
		CatID: req.CatID,
	}
}
