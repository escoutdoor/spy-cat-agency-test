package mission

import (
	"context"

	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

func (s *service) ListMissions(ctx context.Context, limit, offset int) ([]entity.Mission, error) {
	missions, err := s.missionRepository.ListMissions(ctx, limit, offset)
	if err != nil {
		return nil, errwrap.Wrap("list missions from repository", err)
	}

	return missions, nil
}
