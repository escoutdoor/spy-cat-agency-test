package v1

import (
	"errors"

	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
	"github.com/gofiber/fiber/v3"
)

func (c *controller) deleteTarget(ctx fiber.Ctx) error {
	targetID := ctx.Params(targetIDParams)
	if err := validateTargetID(targetID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.targetService.DeleteTarget(ctx, targetID); err != nil {
		appErr := new(apperrors.Error)
		if errors.As(err, &appErr) {
			switch appErr.Code {
			case apperrors.TargetNotFound:
				return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": appErr.Error(),
				})
			case apperrors.TargetAlreadyCompleted:
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": appErr.Error(),
				})
			}
		}

		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
