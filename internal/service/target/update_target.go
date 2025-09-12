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

	if txErr := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		if err := s.targetRepository.UpdateTarget(ctx, in); err != nil {
			return errwrap.Wrap("update target", err)
		}

		count, err := s.targetRepository.CountIncompliteTargets(ctx, mission.ID)
		if err != nil {
			return errwrap.Wrap("count incomplite targets", err)
		}

		if count == 0 {
			completed := true
			params := dto.UpdateMissionParams{
				ID:        mission.ID,
				Completed: &completed,
			}

			if err := s.missionRepository.UpdateMission(ctx, params); err != nil {
				return errwrap.Wrap("update mission", err)
			}
		}

		return nil
	}); txErr != nil {
		return txErr
	}

	return nil
}
