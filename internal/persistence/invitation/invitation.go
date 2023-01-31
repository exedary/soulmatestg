package invitation

import (
	"context"

	"github.com/exedary/soulmates/internal/domain/invitation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName         = "soulmates"
	collectionName = "invitations"
)

type Repository struct {
	client      *mongo.Client
	invitations *mongo.Collection
}

func NewRepository(client *mongo.Client) *Repository {
	client.Database(dbName).CreateCollection(context.TODO(), collectionName)
	return &Repository{
		client:      client,
		invitations: client.Database(dbName).Collection(collectionName),
	}
}

func (repository *Repository) Create(ctx context.Context, invitation *invitation.Invitation) error {
	if _, err := repository.invitations.InsertOne(ctx, invitation); err != nil {
		return err
	}

	return nil
}

func (repository *Repository) GetById(ctx context.Context, id string) (*invitation.Invitation, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	invitation := &invitation.Invitation{}

	if err := repository.invitations.FindOne(ctx, byIdSpec(objectId)).Decode(invitation); err != nil {
		return nil, err
	}

	return invitation, nil
}
