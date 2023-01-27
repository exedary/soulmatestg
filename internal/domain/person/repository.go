package person

import "context"

type Repository interface {
	GetById(ctx context.Context, id string) (*Person, error)
	GetByExternalId(ctx context.Context, externalId string) (*Person, error)
	Create(ctx context.Context, person *Person) (string, error)
}
