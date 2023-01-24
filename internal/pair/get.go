package pair

import (
	"context"
	"time"

	"github.com/exedary/soulmates/internal/domain/pair"
)

type pairResoponseDto struct {
	id          string
	dateCreated string
}

func GetById(ctx context.Context, repository pair.Repository, id string) (*pairResoponseDto, error) {
	pair, err := repository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &pairResoponseDto{
		pair.Id.Hex(),
		pair.DateCreated.Format(time.UnixDate),
	}, nil
}
