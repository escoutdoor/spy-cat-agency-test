package cat

import (
	"context"

	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

func (s *service) GetCat(ctx context.Context, catID string) (entity.Cat, error) {
	cat, err := s.catRepository.GetCat(ctx, catID)
	if err != nil {
		return entity.Cat{}, errwrap.Wrap("get cat from repository", err)
	}

	return cat, nil
}
