package cat

import (
	"github.com/escoutdoor/spy-cat-agency-test/internal/client"
	"github.com/escoutdoor/spy-cat-agency-test/internal/repository"
)

type service struct {
	catRepository repository.CatRepisitory
	catApiClient  client.CatClient
}

func New(catRepository repository.CatRepisitory, catApiClient client.CatClient) *service {
	return &service{
		catRepository: catRepository,
		catApiClient:  catApiClient,
	}
}
