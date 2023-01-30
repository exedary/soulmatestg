package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(func(*http.Server) {}),
)

func New(lc fx.Lifecycle, router *gin.Engine, logger *zap.SugaredLogger) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	lc.Append(fx.StartHook(func(ctx context.Context) {
		logger.Info("Listening on localhost:8080")
		go srv.ListenAndServe()
	}))

	lc.Append(fx.StopHook(func(ctx context.Context) error {
		logger.Info("Stopping server")
		return srv.Shutdown(ctx)
	}))

	return srv
}
