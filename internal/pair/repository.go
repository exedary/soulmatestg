package pair

import (
	"context"
	"errors"

	"github.com/exedary/soulmates/internal/entity"
	"github.com/exedary/soulmates/internal/uuid"
)

type Repository interface {
	Create(ctx context.Context, pair entity.Pair) (string, error)
	GetById(ctx context.Context, id string) (*entity.Pair, error)
}

func NewRepo() Repository {
	return &repository{
		storage: make(map[string]entity.Pair),
	}
}

type repository struct {
	storage map[string]entity.Pair
}

func (repository repository) GetById(ctx context.Context, id string) (*entity.Pair, error) {
	ok, err := repository.storage[id]

	if err {
		return &ok, nil
	}

	return nil, errors.New("not found")
}

func (repository repository) Create(ctx context.Context, pair entity.Pair) (string, error) {
	pair.Id = uuid.Generate()
	repository.storage[pair.Id] = pair

	return pair.Id, nil
}
