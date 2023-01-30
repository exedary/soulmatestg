package main

import (
	"github.com/exedary/soulmates/internal/auth"
	"github.com/exedary/soulmates/internal/config"
	"github.com/exedary/soulmates/internal/logger"
	"github.com/exedary/soulmates/internal/mongo"
	"github.com/exedary/soulmates/internal/pair"
	"github.com/exedary/soulmates/internal/persistence"
	"github.com/exedary/soulmates/internal/router"
	"github.com/exedary/soulmates/internal/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		router.Module,
		server.Module,
		logger.Module,
		mongo.Module,
		persistence.Module,
		pair.Module,
		auth.Module,
	).Run()
}
