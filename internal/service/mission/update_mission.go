package mission

import (
	"context"

	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

func (s *service) UpdateMission(ctx context.Context, in dto.UpdateMissionParams) error {
	if _, err := s.missionRepository.GetMission(ctx, in.ID); err != nil {
		return errwrap.Wrap("get mission from repository", err)
	}

	if in.CatID != nil {
		if _, err := s.catRepository.GetCat(ctx, *in.CatID); err != nil {
			return errwrap.Wrap("get cat from repository", err)
		}

		onMission, err := s.missionRepository.IsCatOnMission(ctx, *in.CatID)
		if err != nil {
			return errwrap.Wrap("check is cat on the mission", err)
		}

		if onMission {
			return apperrors.CatOnMissionWithID(*in.CatID)
		}
	}

	if err := s.missionRepository.UpdateMission(ctx, in); err != nil {
		return errwrap.Wrap("update mission", err)
	}

	return nil
}
