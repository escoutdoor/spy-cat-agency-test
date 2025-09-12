package cat

import (
	"context"

	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

func (s *service) DeleteCat(ctx context.Context, catID string) error {
	if _, err := s.catRepository.GetCat(ctx, catID); err != nil {
		return errwrap.Wrap("get cat from repository", err)
	}

	if err := s.catRepository.DeleteCat(ctx, catID); err != nil {
		return errwrap.Wrap("delete cat", err)
	}

	return nil
}
