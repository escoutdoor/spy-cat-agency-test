package v1

import (
	"errors"

	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type addTargetsRequestBody struct {
	Targets []struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"targets" validate:"required,min=1,max=3"`
}

func (c *controller) addTargets(ctx fiber.Ctx) error {
	missionID := ctx.Params(missionIDParam)
	if err := validateMissionID(missionID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var req addTargetsRequestBody

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

	if err := c.targetService.AddTargets(ctx, missionID, addTargetsRequestBodyToCreateMissionParams(req)); err != nil {
		appErr := new(apperrors.Error)
		if errors.As(err, &appErr) {
			switch appErr.Code {
			case apperrors.MissionAlreadyCompleted:
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": appErr.Error(),
				})
			case apperrors.MissionNotFound:
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": appErr.Error(),
				})
			case apperrors.TargetLimit:
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": appErr.Error(),
				})
			}
		}

		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func addTargetsRequestBodyToCreateMissionParams(req addTargetsRequestBody) []dto.CreateTargetParams {
	targets := make([]dto.CreateTargetParams, 0, len(req.Targets))
	for _, t := range req.Targets {
		targets = append(targets, dto.CreateTargetParams{
			Name:    t.Name,
			Country: t.Country,
		})
	}

	return targets
}
