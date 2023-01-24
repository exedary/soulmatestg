package pair

import "context"

type Repository interface {
	GetById(ctx context.Context, id string) (*Pair, error)
	Create(ctx context.Context, pair *Pair) (string, error)
}
