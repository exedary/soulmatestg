package router

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(provideGinEngine),
	fx.Provide(provideV1),
)

func provideGinEngine(logger *zap.Logger) *gin.Engine {
	router := gin.New()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	return router
}

func provideV1(gin *gin.Engine) *gin.RouterGroup {
	return gin.Group("api/v1")
}
