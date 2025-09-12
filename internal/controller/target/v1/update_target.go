package v1

import (
	"errors"

	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type updateTargetRequestBody struct {
	Completed *bool   `json:"completed"`
	Notes     *string `json:"catId"`
}

func (c *controller) updateTarget(ctx fiber.Ctx) error {
	targetID := ctx.Params(targetIDParams)
	if err := validateTargetID(targetID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var req updateTargetRequestBody
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

	if err := c.targetService.UpdateTarget(ctx, updateTargetRequestBodyToUpdateTargetParams(req, targetID)); err != nil {
		appErr := new(apperrors.Error)
		if errors.As(err, &appErr) {
			switch appErr.Code {
			case apperrors.MissionAlreadyCompleted:
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": appErr.Error(),
				})
			case apperrors.NoFieldsNeedToBeUpdated:
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": appErr.Error(),
				})
			case apperrors.TargetAlreadyCompleted:
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": appErr.Error(),
				})
			case apperrors.TargetNotFound:
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": appErr.Error(),
				})
			}
		}

		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func updateTargetRequestBodyToUpdateTargetParams(req updateTargetRequestBody, targetID string) dto.UpdateTargetParams {
	return dto.UpdateTargetParams{
		ID:        targetID,
		Completed: req.Completed,
		Notes:     req.Notes,
	}
}
