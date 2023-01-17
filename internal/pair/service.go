package pair

import "context"

type IService interface {
	Create(ctx context.Context, pairModel *CreatePairDto) (string, error)
	GetById(ctx context.Context, id string) (PairDto, error)
}

type service struct {
	Repository *Repository
}
