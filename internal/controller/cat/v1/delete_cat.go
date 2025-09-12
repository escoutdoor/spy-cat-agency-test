package v1

import (
	"errors"

	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
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
		appErr := new(apperrors.Error)
		if errors.As(err, &appErr) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": appErr.Error(),
			})
		}

		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
