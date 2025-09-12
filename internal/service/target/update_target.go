package target

import (
	"context"

	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

func (s *service) UpdateTarget(ctx context.Context, in dto.UpdateTargetParams) error {
	target, err := s.targetRepository.GetTarget(ctx, in.ID)
	if err != nil {
		return errwrap.Wrap("get target from repository", err)
	}
	mission, err := s.missionRepository.GetMission(ctx, target.MissionID)
	if err != nil {
		return errwrap.Wrap("get mission from repository", err)
	}

	if target.Completed {
		return apperrors.TargetAlreadyCompletedWithID(target.ID)
	}
	if mission.Completed {
		return apperrors.MissionAlreadyCompletedWithID(target.MissionID)
	}

	if err := s.targetRepository.UpdateTarget(ctx, in); err != nil {
		return errwrap.Wrap("update target", err)
	}

	return nil
}
