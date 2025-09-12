package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func (c *controller) listMissions(ctx fiber.Ctx) error {
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

	missions, err := c.missionService.ListMissions(ctx, limit, offset)
	if err != nil {
		return err
	}

	type response struct {
		Missions []missionResponse `json:"missions"`
	}

	missionsResponse := make([]missionResponse, 0, len(missions))
	for _, m := range missions {
		missionsResponse = append(missionsResponse, missionToMissionResponse(m))
	}

	return ctx.Status(fiber.StatusOK).JSON(response{Missions: missionsResponse})
}
