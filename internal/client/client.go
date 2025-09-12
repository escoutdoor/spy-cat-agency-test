package client

import "context"

type CatClient interface {
	Exists(ctx context.Context, breed string) (bool, error)
}
