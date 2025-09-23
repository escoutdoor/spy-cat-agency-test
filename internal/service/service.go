package service

import (
	"context"

	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
)

type CatService interface {
	GetCat(ctx context.Context, catID string) (entity.Cat, error)
	ListCats(ctx context.Context, limit, offset int) ([]entity.Cat, error)
	UpdateCat(ctx context.Context, in dto.UpdateCatParams) (entity.Cat, error)
	DeleteCat(ctx context.Context, catID string) error
	CreateCat(ctx context.Context, in dto.CreateCatParams) (entity.Cat, error)
}

type MissionService interface {
	CreateMission(ctx context.Context, in dto.CreateMissionParams) (string, error)
	DeleteMission(ctx context.Context, missionID string) error
	UpdateMission(ctx context.Context, in dto.UpdateMissionParams) error
	ListMissions(ctx context.Context, limit, offset int) ([]entity.Mission, error)
	GetMission(ctx context.Context, missionID string) (entity.Mission, error)
	AddTargets(ctx context.Context, missionID string, in []dto.CreateTargetParams) error
}

type TargetService interface {
	DeleteTarget(ctx context.Context, targetID string) error
	UpdateTarget(ctx context.Context, in dto.UpdateTargetParams) error
}
