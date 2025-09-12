package v1

import (
	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type createMissionRequestBody struct {
	Targets []struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"targets" validate:"required,min=1,max=3"`
}

func (c *controller) createMission(ctx fiber.Ctx) error {
	var req createMissionRequestBody

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

	missionID, err := c.missionService.CreateMission(ctx, createMissionRequestBodyToCreateMissionParams(req))
	if err != nil {
		return err
	}

	type response struct {
		MissionID string `json:"missionId"`
	}

	return ctx.Status(fiber.StatusCreated).JSON(response{MissionID: missionID})
}

func createMissionRequestBodyToCreateMissionParams(req createMissionRequestBody) dto.CreateMissionParams {
	targets := make([]dto.CreateTargetParams, 0, len(req.Targets))
	for _, t := range req.Targets {
		targets = append(targets, dto.CreateTargetParams{
			Name:    t.Name,
			Country: t.Country,
		})
	}

	return dto.CreateMissionParams{
		Targets: targets,
	}
}
