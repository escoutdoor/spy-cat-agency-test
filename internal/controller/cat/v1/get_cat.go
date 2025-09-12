package v1

import (
	"errors"

	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
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
		appErr := new(apperrors.Error)
		if errors.As(err, &appErr) {
			if appErr.Code == apperrors.CatNotFound {
				return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": appErr.Error(),
				})
			}
		}

		return err
	}

	type response struct {
		Cat entity.Cat `json:"cat"`
	}
	return ctx.Status(fiber.StatusOK).JSON(response{Cat: cat})
}
