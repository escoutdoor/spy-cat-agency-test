package repository

import (
	"context"

	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
)

type CatRepisitory interface {
	GetCat(ctx context.Context, catID string) (entity.Cat, error)
	ListCats(ctx context.Context, limit, offset int) ([]entity.Cat, error)
	UpdateCat(ctx context.Context, in dto.UpdateCatParams) (entity.Cat, error)
	DeleteCat(ctx context.Context, catID string) error
	CreateCat(ctx context.Context, in dto.CreateCatParams) (entity.Cat, error)
}

type MissionRepository interface {
	GetMission(ctx context.Context, missionID string) (entity.Mission, error)
	ListMissions(ctx context.Context, limit, offset int) ([]entity.Mission, error)
	UpdateMission(ctx context.Context, in dto.UpdateMissionParams) error
	DeleteMission(ctx context.Context, missionID string) error
	CreateMission(ctx context.Context, in dto.CreateMissionParams) (string, error)
	IsCatOnMission(ctx context.Context, catID string) (bool, error)
}

type TargetRepository interface {
	GetTarget(ctx context.Context, targetID string) (entity.Target, error)
	UpdateTarget(ctx context.Context, in dto.UpdateTargetParams) error
	DeleteTarget(ctx context.Context, targetID string) error
	CreateTarget(ctx context.Context, missionID string, in dto.CreateTargetParams) (string, error)
	CountIncompliteTargets(ctx context.Context, missionID string) (int, error)
}
