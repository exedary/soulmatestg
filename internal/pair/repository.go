package pair

import (
	"context"

	"github.com/exedary/soulmates/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, pair entity.Pair) error
	GetById(ctx context.Context, id string) (*entity.Pair, error)
}
