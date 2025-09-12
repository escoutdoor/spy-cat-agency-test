package v1

import (
	"errors"

	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type createCatRequestBody struct {
	Name              string  `json:"name" validate:"required,min=3"`
	YearsOfExperience int     `json:"yearsOfExperience" validate:"required,gte=0"`
	Breed             string  `json:"breed" validate:"required,min=2"`
	Salary            float64 `json:"salary" validate:"required,gt=1"`
}

func (c *controller) createCat(ctx fiber.Ctx) error {
	var req createCatRequestBody

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

	cat, err := c.catService.CreateCat(ctx, createCatRequestBodyToCreateCatParams(req))
	if err != nil {
		appErr := new(apperrors.Error)
		if errors.As(err, &appErr) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": appErr.Error(),
			})
		}
		return err
	}

	type response struct {
		Cat catResponse `json:"cat"`
	}
	return ctx.Status(fiber.StatusCreated).JSON(response{Cat: catToCatResponse(cat)})
}

func createCatRequestBodyToCreateCatParams(req createCatRequestBody) dto.CreateCatParams {
	return dto.CreateCatParams{
		Name:              req.Name,
		YearsOfExperience: req.YearsOfExperience,
		Breed:             req.Breed,
		Salary:            req.Salary,
	}
}
