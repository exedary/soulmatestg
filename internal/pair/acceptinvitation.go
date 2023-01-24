package pair

import (
	"context"

	"github.com/exedary/soulmates/internal/domain/pair"
)

func AcceptInvitation(ctx context.Context, repository pair.Repository) error {
	return ctx.Err()
}
