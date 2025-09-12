package cat

import (
	"context"

	"github.com/escoutdoor/spy-cat-agency-test/internal/dto"
	"github.com/escoutdoor/spy-cat-agency-test/internal/entity"
	apperrors "github.com/escoutdoor/spy-cat-agency-test/internal/errors"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
)

func (s *service) CreateCat(ctx context.Context, in dto.CreateCatParams) (entity.Cat, error) {
	exists, err := s.catApiClient.Exists(ctx, in.Breed)
	if err != nil {
		return entity.Cat{}, errwrap.Wrap("check if cat exists", err)
	}
	if !exists {
		return entity.Cat{}, apperrors.BreedDoesNotExist
	}

	cat, err := s.catRepository.CreateCat(ctx, in)
	if err != nil {
		return entity.Cat{}, errwrap.Wrap("create cat", err)
	}

	return cat, nil
}
