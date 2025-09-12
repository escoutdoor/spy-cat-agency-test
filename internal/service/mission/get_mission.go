package mission

import (
	"context"

	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

func (s *service) GetMission(ctx context.Context, missionID string) (entity.Mission, error) {
	mission, err := s.missionRepository.GetMission(ctx, missionID)
	if err != nil {
		return entity.Mission{}, errwrap.Wrap("get mission from repository", err)
	}

	return mission, nil
}
