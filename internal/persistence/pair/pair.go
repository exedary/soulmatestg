package persistence

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

type PairRepository struct {
	client *mongo.Client
	pairs  *mongo.Collection
}

func NewPairRepository(client *mongo.Client) *PairRepository {
	client.Database(dbName).CreateCollection(context.TODO(), collectionName)
	return &PairRepository{
		client: client,
		pairs:  client.Database(dbName).Collection(collectionName),
	}
}

func (repository *PairRepository) Create(ctx context.Context, pair *pair.Pair) (string, error) {
	if _, err := repository.pairs.InsertOne(ctx, pair); err != nil {
		return "", err
	}

	return pair.Id.Hex(), nil
}

func (repository *PairRepository) GetById(ctx context.Context, id string) (*pair.Pair, error) {
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
