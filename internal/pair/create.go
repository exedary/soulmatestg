package pair

import (
	"context"
	"errors"

	"github.com/exedary/soulmates/internal/domain/invitation"
	"github.com/exedary/soulmates/internal/domain/pair"
	"github.com/exedary/soulmates/internal/domain/person"
)

type entityIdDto struct {
	id string
}

func Create(ctx context.Context, pairRepository pair.Repository, invitationRepository invitation.Repository, personId string) (*entityIdDto, error) {
	oldPair, err := pairRepository.FindPairByPerson(ctx, personId)

	if err != nil {
		return nil, err
	}

	if oldPair != nil {
		return nil, errors.New("pair with this person already exists")
	}

	newPair := pair.New(&person.Person{})

	id, err := pairRepository.Create(ctx, newPair)

	if err != nil {
		return nil, err
	}

	return &entityIdDto{id: id}, nil
}
