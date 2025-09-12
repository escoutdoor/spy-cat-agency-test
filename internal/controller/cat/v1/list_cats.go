package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func (c *controller) listCats(ctx fiber.Ctx) error {
	var (
		limit  = 0
		offset = 0
		err    error
	)

	limitStr := ctx.Query("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "limit value should be an integer")
		}
	}

	offsetStr := ctx.Query("offset")
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "offset value should be an integer")
		}
	}

	cats, err := c.catService.ListCats(ctx, limit, offset)
	if err != nil {
		return err
	}

	type response struct {
		Cats []catResponse `json:"cats"`
	}

	respCat := make([]catResponse, 0, len(cats))
	for _, c := range cats {
		respCat = append(respCat, catToCatResponse(c))
	}

	return ctx.Status(fiber.StatusOK).JSON(response{Cats: respCat})
}
