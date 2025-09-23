package v1

import (
	"fmt"

	"github.com/escoutdoor/spy-cat-agency-test/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

const (
	targetIDParams = "targetId"
)

type controller struct {
	targetService service.TargetService
}

func Register(a *fiber.App, targetService service.TargetService) {
	ctl := &controller{targetService: targetService}
	r := a.Group("/v1/targets")

	r.Delete("/:targetId", ctl.deleteTarget)
	r.Patch("/:targetId", ctl.updateTarget)
}

func validateTargetID(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("targetId parameter is required")
	}

	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("invalid targetId parameter, should be uuid")
	}

	return nil
}
