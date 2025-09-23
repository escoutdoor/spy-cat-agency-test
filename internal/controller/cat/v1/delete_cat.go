package v1

import (
	"github.com/gofiber/fiber/v3"
)

func (c *controller) deleteCat(ctx fiber.Ctx) error {
	catID := ctx.Params(catIDParam)
	if err := validateCatID(catID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.catService.DeleteCat(ctx, catID); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
