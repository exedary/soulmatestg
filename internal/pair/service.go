package pair

import (
	"context"
	"time"

	"github.com/exedary/soulmates/internal/entity"
)

type Service interface {
	Create(ctx context.Context, pairModel *createPairDto) (string, error)
	GetById(ctx context.Context, id string) (*pairDto, error)
}

type service struct {
	Repository Repository
}

func New(repository Repository) Service {
	return service{
		Repository: repository,
	}
}

func (service service) Create(ctx context.Context, pairModel *createPairDto) (string, error) {
	invitedUser := entity.User{
		PhoneNumber: pairModel.InvitedId,
	}

	pair := entity.Pair{
		DateCreated: time.Now(),

		Acceptance: entity.PairAcceptance{User: invitedUser, IsAccepted: false},
	}
	id, err := service.Repository.Create(ctx, pair)

	if err != nil {
		return "", err
	}

	return id, nil

}

func (service service) GetById(ctx context.Context, id string) (*pairDto, error) {
	pair, err := service.Repository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &pairDto{
		pair.Id,
		pair.DateCreated.Format(time.RFC3339),
	}, nil
}
