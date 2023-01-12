package runner

import (
	configs "github.com/exedary/soulmates/pkg/configuration"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	defaultSwaggerRoute = "/swagger/*any"
)

type Gin struct {
	configuration *configs.ConfigurationBase
	GinEngine     *gin.Engine
}

func NewGinRunner() *Gin {
	ginRunner := &Gin{GinEngine: gin.Default(), configuration: &configs.ConfigurationBase{}}
	return ginRunner
}

func (gin *Gin) UseConfiguration(configPath string) *Gin {
	configration, err := configs.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	gin.configuration = configration

	return gin
}

func (gin *Gin) UseSwagger() *Gin {
	gin.GinEngine.GET(defaultSwaggerRoute, ginSwagger.WrapHandler(swaggerfiles.Handler))

	return gin
}

func (gin *Gin) Run() {
	gin.GinEngine.Run(gin.configuration.Port)
}
