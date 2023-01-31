package person

import (
	"context"

	"github.com/exedary/soulmates/internal/domain/person"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

const (
	dbName         = "soulmates"
	collectionName = "persons"
)

var Module = fx.Provide(NewRepository)

type Repository struct {
	client  *mongo.Client
	persons *mongo.Collection
}

func NewRepository(client *mongo.Client) *Repository {
	client.Database(dbName).CreateCollection(context.TODO(), collectionName)
	return &Repository{
		client:  client,
		persons: client.Database(dbName).Collection(collectionName),
	}
}

func (repository *Repository) Create(ctx context.Context, person *person.Person) (string, error) {
	if _, err := repository.persons.InsertOne(ctx, person); err != nil {
		return "", err
	}

	return person.Id.Hex(), nil
}

func (repository *Repository) GetById(ctx context.Context, id string) (*person.Person, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	person := person.Person{}

	if err := repository.persons.FindOne(ctx, byIdSpec(objectId)).Decode(person); err != nil {
		return nil, err
	}

	return &person, nil
}

func (repository *Repository) GetByExternalId(ctx context.Context, externalId string) (*person.Person, error) {
	person := person.Person{}

	if err := repository.persons.FindOne(ctx, byExternalIdSpec(externalId)).Decode(person); err != nil {
		return nil, err
	}

	return &person, nil
}
