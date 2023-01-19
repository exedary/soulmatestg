package main

import (
	docs "github.com/exedary/soulmates/docs"
	"github.com/exedary/soulmates/internal/pair"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {

}

func main() {
	app := gin.Default()

	v1 := app.Group("api/v1")
	docs.SwaggerInfo.BasePath = "/api/v1"
	pairService := pair.New(pair.NewRepo())

	pair.RegisterHandlers(v1, &pairService)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.Run(":8080")
}
