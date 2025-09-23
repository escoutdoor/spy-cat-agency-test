package v1

import (
	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	"github.com/gofiber/fiber/v3"
)

func (c *controller) getCat(ctx fiber.Ctx) error {
	catID := ctx.Params(catIDParam)
	if err := validateCatID(catID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cat, err := c.catService.GetCat(ctx, catID)
	if err != nil {
		return err
	}

	type response struct {
		Cat entity.Cat `json:"cat"`
	}
	return ctx.Status(fiber.StatusOK).JSON(response{Cat: cat})
}
