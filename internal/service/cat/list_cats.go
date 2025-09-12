package cat

import (
	"context"

	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

func (s *service) ListCats(ctx context.Context, limit, offset int) ([]entity.Cat, error) {
	cats, err := s.catRepository.ListCats(ctx, limit, offset)
	if err != nil {
		return nil, errwrap.Wrap("list cats from repository", err)
	}

	return cats, nil
}
