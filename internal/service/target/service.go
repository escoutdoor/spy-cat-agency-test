package target

import (
	"github.com/escoutdoor/spy-cat-agency-test/internal/repository"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/database"
)

type service struct {
	missionRepository repository.MissionRepository
	targetRepository  repository.TargetRepository
	txManager         database.TxManager
}

func New(
	missionRepository repository.MissionRepository,
	targetRepository repository.TargetRepository,
	txManager database.TxManager,
) *service {
	return &service{
		missionRepository: missionRepository,
		targetRepository:  targetRepository,
		txManager:         txManager,
	}
}
