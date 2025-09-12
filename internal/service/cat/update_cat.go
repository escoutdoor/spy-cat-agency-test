package cat

import (
	"context"

	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

func (s *service) UpdateCat(ctx context.Context, in dto.UpdateCatParams) (entity.Cat, error) {
	if _, err := s.catRepository.GetCat(ctx, in.ID); err != nil {
		return entity.Cat{}, errwrap.Wrap("get cat from repository", err)
	}

	cat, err := s.catRepository.UpdateCat(ctx, in)
	if err != nil {
		return entity.Cat{}, errwrap.Wrap("update cat", err)
	}

	return cat, nil
}
