package target

import (
	"context"

	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

func (s *service) DeleteTarget(ctx context.Context, targetID string) error {
	target, err := s.targetRepository.GetTarget(ctx, targetID)
	if err != nil {
		return errwrap.Wrap("get target from repository", err)
	}

	if target.Completed {
		return apperrors.TargetAlreadyCompletedWithID(targetID)
	}

	if err := s.targetRepository.DeleteTarget(ctx, targetID); err != nil {
		return errwrap.Wrap("delete target", err)
	}
	return nil
}
