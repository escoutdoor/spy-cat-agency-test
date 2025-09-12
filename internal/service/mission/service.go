package mission

import (
	"github.com/escoutdoor/spy-cat-agency-test/internal/repository"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/database"
)

type service struct {
	catRepository     repository.CatRepisitory
	missionRepository repository.MissionRepository
	targetRepository  repository.TargetRepository
	txManager         database.TxManager
}

func New(
	catRepository repository.CatRepisitory,
	missionRepository repository.MissionRepository,
	targetRepository repository.TargetRepository,
	txManager database.TxManager,
) *service {
	return &service{
		catRepository:     catRepository,
		missionRepository: missionRepository,
		targetRepository:  targetRepository,
		txManager:         txManager,
	}
}
