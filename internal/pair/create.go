package pair

import (
	"context"

	"github.com/exedary/soulmates/internal/domain/pair"
	"github.com/exedary/soulmates/internal/domain/person"
)

type entityIdDto struct {
	id string
}

func Create(ctx context.Context, pairRepository pair.Repository) (*entityIdDto, error) {
	pair := pair.New(&person.Person{})

	id, err := pairRepository.Create(ctx, pair)

	if err != nil {
		return nil, err
	}

	return &entityIdDto{id: id}, nil
}
