package mongo

import (
	"context"
	"log"
	"time"

	"github.com/exedary/soulmates/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(provideMongo)

const defaultTimeout = 10 * time.Second

func provideMongo(lc fx.Lifecycle, config *config.Configuration, logger *zap.SugaredLogger) (*mongo.Client, error) {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(config.Database.Uri),
		/*options.Client().SetAuth(options.Credential{
			Username: "",
			Password: "",
		})*/)

	if err != nil {
		log.Fatal(err)
	}

	lc.Append(fx.StartHook(func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
		defer cancel()

		logger.Info("Connecting to mongo db")

		return client.Connect(ctx)
	}))

	lc.Append(fx.StopHook(func(ctx context.Context) error {
		return client.Disconnect(ctx)
	}))

	return client, nil
}
