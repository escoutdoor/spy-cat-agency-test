package mission

import (
	"context"
	"fmt"

	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

func (s *service) AddTargets(ctx context.Context, missionID string, in []dto.CreateTargetParams) error {
	mission, err := s.missionRepository.GetMission(ctx, missionID)
	if err != nil {
		return errwrap.Wrap("get mission from repository", err)
	}
	if mission.Completed {
		return apperrors.MissionAlreadyCompletedWithID(missionID)
	}
	if len(mission.Targets)+len(in) > 3 {
		return apperrors.TargetsLimitErr
	}

	if txErr := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		for _, t := range in {
			// TODO: create targets with multiple rows to create
			if _, err := s.targetRepository.CreateTarget(ctx, missionID, t); err != nil {
				msg := fmt.Sprintf("add target with name %q to mission %q", t.Name, missionID)
				return errwrap.Wrap(msg, err)
			}
		}

		return nil
	}); txErr != nil {
		return err
	}

	return nil
}
