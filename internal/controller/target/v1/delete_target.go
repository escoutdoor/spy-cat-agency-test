package v1

import (
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
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
