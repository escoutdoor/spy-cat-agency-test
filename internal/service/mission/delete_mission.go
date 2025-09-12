package mission

import (
	"context"

	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

func (s *service) DeleteMission(ctx context.Context, missionID string) error {
	mission, err := s.missionRepository.GetMission(ctx, missionID)
	if err != nil {
		return errwrap.Wrap("get mission from repository", err)
	}
	if mission.CatID != nil {
		return apperrors.MissionCannotBeDeleted(missionID)
	}

	if err := s.missionRepository.DeleteMission(ctx, missionID); err != nil {
		return errwrap.Wrap("delete mission", err)
	}

	return nil
}
