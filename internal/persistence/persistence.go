package persistence

import (
	domainPair "github.com/exedary/soulmates/internal/domain/pair"
	domainPerson "github.com/exedary/soulmates/internal/domain/person"
	"github.com/exedary/soulmates/internal/persistence/pair"
	"github.com/exedary/soulmates/internal/persistence/person"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		person.NewRepository,
		fx.As(new(domainPerson.Repository))),
	fx.Annotate(
		pair.NewRepository,
		fx.As(new(domainPair.Repository))),
)
