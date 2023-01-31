package pair

import (
	"context"

	"github.com/exedary/soulmates/internal/domain/pair"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName         = "soulmates"
	collectionName = "pairs"
)

type Repository struct {
	client *mongo.Client
	pairs  *mongo.Collection
}

func NewRepository(client *mongo.Client) *Repository {
	client.Database(dbName).CreateCollection(context.TODO(), collectionName)
	return &Repository{
		client: client,
		pairs:  client.Database(dbName).Collection(collectionName),
	}
}

func (repository *Repository) Create(ctx context.Context, pair *pair.Pair) (string, error) {
	if _, err := repository.pairs.InsertOne(ctx, pair); err != nil {
		return "", err
	}

	return pair.Id.Hex(), nil
}

func (repository *Repository) GetById(ctx context.Context, id string) (*pair.Pair, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	pair := pair.Pair{}

	if err := repository.pairs.FindOne(ctx, byIdSpec(objectId)).Decode(pair); err != nil {
		return nil, err
	}

	return &pair, nil
}

func (repository *Repository) FindPairByPerson(ctx context.Context, personId string) (*pair.Pair, error) {
	objectId, err := primitive.ObjectIDFromHex(personId)

	if err != nil {
		return nil, err
	}

	pair := pair.Pair{}

	if err := repository.pairs.FindOne(ctx, byPersonIdSpec(objectId)).Decode(pair); err != nil {
		return nil, err
	}

	return &pair, nil
}
