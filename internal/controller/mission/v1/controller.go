package v1

import (
	"fmt"
	"time"

	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	"github.com/escoutdoor/spy-cat-agency-test/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

const (
	missionIDParam = "missionId"
)

type controller struct {
	missionService service.MissionService
}

func Register(a *fiber.App, missionService service.MissionService) {
	ctl := &controller{missionService: missionService}
	r := a.Group("/v1/missions")

	r.Post("/", ctl.createMission)
	r.Delete("/:missionId", ctl.deleteMission)
	r.Get("/:missionId", ctl.getMission)
	r.Get("/", ctl.listMissions)
	r.Patch("/:missionId", ctl.updateMission)
}

func validateMissionID(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("id parameter is required")
	}

	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("invalid id parameter, should be uuid")
	}

	return nil
}

type targetResponse struct {
	ID        string    `json:"id"`
	MissionID string    `json:"missionId"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	Notes     string    `json:"notes"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updatedAt"`
}

type missionResponse struct {
	ID        string           `json:"id"`
	CatID     *string          `json:"catId,omitempty"`
	Targets   []targetResponse `json:"targets"`
	Completed bool             `json:"completed"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}

func targetToTargetResponse(target entity.Target) targetResponse {
	return targetResponse{
		ID:        target.ID,
		MissionID: target.MissionID,
		Name:      target.Name,
		Country:   target.Country,
		Notes:     target.Notes,
		Completed: target.Completed,
		CreatedAt: target.CreatedAt,
		UpdateAt:  target.UpdatedAt,
	}
}

func missionToMissionResponse(mission entity.Mission) missionResponse {
	targets := make([]targetResponse, 0, len(mission.Targets))
	for _, t := range mission.Targets {
		targets = append(targets, targetToTargetResponse(t))
	}

	return missionResponse{
		ID:        mission.ID,
		CatID:     mission.CatID,
		Targets:   targets,
		Completed: mission.Completed,
		CreatedAt: mission.CreatedAt,
		UpdatedAt: mission.UpdatedAt,
	}
}
