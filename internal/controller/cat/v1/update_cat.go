package v1

import (
	"errors"

	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type updateCatRequestBody struct {
	Salary float64 `json:"salary" validate:"required,gt=0"`
}

func (c *controller) updateCat(ctx fiber.Ctx) error {
	catID := ctx.Params(catIDParam)
	if err := validateCatID(catID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var req updateCatRequestBody
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

	cat, err := c.catService.UpdateCat(ctx, updateCatRequestBodyToUpdateCatParams(req, catID))
	if err != nil {
		appErr := new(apperrors.Error)
		if errors.As(err, &appErr) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": appErr.Error(),
			})
		}
	}

	type response struct {
		Cat catResponse `json:"cat"`
	}

	return ctx.Status(fiber.StatusOK).JSON(response{Cat: catToCatResponse(cat)})
}

func updateCatRequestBodyToUpdateCatParams(req updateCatRequestBody, catID string) dto.UpdateCatParams {
	return dto.UpdateCatParams{
		ID:     catID,
		Salary: req.Salary,
	}
}
