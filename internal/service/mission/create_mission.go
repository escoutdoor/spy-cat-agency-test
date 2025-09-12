package mission

import (
	"context"
	"fmt"

	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

func (s *service) CreateMission(ctx context.Context, in dto.CreateMissionParams) (string, error) {
	var missionID string
	if txErr := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var err error
		missionID, err = s.missionRepository.CreateMission(ctx, in)
		if err != nil {
			return errwrap.Wrap("create mission", err)
		}

		for _, t := range in.Targets {
			// TODO: create targets with multiple rows to create
			if _, err := s.targetRepository.CreateTarget(ctx, missionID, t); err != nil {
				msg := fmt.Sprintf("create target with name %q", t.Name)
				return errwrap.Wrap(msg, err)
			}
		}

		return nil
	}); txErr != nil {
		return "", txErr
	}

	return missionID, nil
}
