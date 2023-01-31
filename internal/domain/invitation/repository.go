package invitation

import "context"

type Repository interface {
	Create(ctx context.Context, invitation *Invitation) error
	GetById(ctx context.Context, id string) (*Invitation, error)
}
